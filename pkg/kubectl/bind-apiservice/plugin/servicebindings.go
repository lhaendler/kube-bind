/*
Copyright 2022 The Kube Bind Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"

	kubebindv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
	"github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1/helpers"
	conditionsapi "github.com/kube-bind/kube-bind/pkg/apis/third_party/conditions/apis/conditions/v1alpha1"
	"github.com/kube-bind/kube-bind/pkg/apis/third_party/conditions/util/conditions"
	bindclient "github.com/kube-bind/kube-bind/pkg/client/clientset/versioned"
)

func (b *BindAPIServiceOptions) createAPIServiceBindings(ctx context.Context, config *rest.Config, request *kubebindv1alpha1.APIServiceExportRequest, secretName string) ([]*kubebindv1alpha1.APIServiceBinding, error) {
	bindClient, err := bindclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	apiextensionsClient, err := apiextensionsclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	var bindings []*kubebindv1alpha1.APIServiceBinding
	for _, resource := range request.Spec.Resources {
		name := resource.Resource + "." + resource.Group
		existing, err := bindClient.KubeBindV1alpha1().APIServiceBindings().Get(ctx, name, metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return nil, err
		} else if err == nil {
			if existing.Spec.KubeconfigSecretRef.Namespace != "kube-bind" || existing.Spec.KubeconfigSecretRef.Name != secretName {
				return nil, fmt.Errorf("found existing APIServiceBinding %s not from this service provider", name)
			}
			fmt.Fprintf(b.Options.IOStreams.ErrOut, "✅ Updating existing APIServiceBinding %s.\n", existing.Name) // nolint: errcheck
			bindings = append(bindings, existing)

			// checking CRD to match the binding
			crd, err := apiextensionsClient.ApiextensionsV1().CustomResourceDefinitions().Get(ctx, resource.Resource+"."+resource.Group, metav1.GetOptions{})
			if err != nil && !apierrors.IsNotFound(err) {
				return nil, err
			} else if err == nil {
				if !helpers.IsOwnedByBinding(existing.Name, existing.UID, crd.OwnerReferences) {
					return nil, fmt.Errorf("CustomResourceDefinition %s exists, but is not owned by kube-bind", crd.Name)
				}
			}
			continue
		}

		var permissionClaims []kubebindv1alpha1.AcceptablePermissionClaim
		for _, c := range resource.PermissionClaims {
			accepted, err := b.promptYesNo(c)
			if err != nil {
				return nil, err
			}

			var state kubebindv1alpha1.AcceptablePermissionClaimState
			if accepted {
				state = kubebindv1alpha1.ClaimAccepted
			} else {
				state = kubebindv1alpha1.ClaimRejected
			}

			permissionClaims = append(permissionClaims, kubebindv1alpha1.AcceptablePermissionClaim{
				PermissionClaim: c,
				State:           state,
			})
		}

		// create new APIServiceBinding.
		first := true
		if err := wait.PollInfinite(1*time.Second, func() (bool, error) {
			if !first {
				first = false
				fmt.Fprint(b.Options.IOStreams.ErrOut, ".") // nolint: errcheck
			}
			created, err := bindClient.KubeBindV1alpha1().APIServiceBindings().Create(ctx, &kubebindv1alpha1.APIServiceBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resource.Resource + "." + resource.Group,
					Namespace: "kube-bind",
				},
				Spec: kubebindv1alpha1.APIServiceBindingSpec{
					KubeconfigSecretRef: kubebindv1alpha1.ClusterSecretKeyRef{
						LocalSecretKeyRef: kubebindv1alpha1.LocalSecretKeyRef{
							Name: secretName,
							Key:  "kubeconfig",
						},
						Namespace: "kube-bind",
					},
					PermissionClaims: permissionClaims,
				},
			}, metav1.CreateOptions{})
			if err != nil {
				return false, err
			}

			// best effort status update to have "Pending" in the Ready condition
			conditions.MarkFalse(created,
				conditionsapi.ReadyCondition,
				"Pending",
				conditionsapi.ConditionSeverityInfo,
				"Pending",
			)
			_, _ = bindClient.KubeBindV1alpha1().APIServiceBindings().UpdateStatus(ctx, created, metav1.UpdateOptions{}) // nolint:errcheck

			fmt.Fprintf(b.Options.IOStreams.ErrOut, "✅ Created APIServiceBinding %s.%s\n", resource.Resource, resource.Group) // nolint: errcheck
			bindings = append(bindings, created)
			return true, nil
		}); err != nil {
			fmt.Fprintln(b.Options.IOStreams.ErrOut, "") // nolint: errcheck
			return nil, err
		}
	}

	return bindings, nil
}

func (opt BindAPIServiceOptions) promptYesNo(p kubebindv1alpha1.PermissionClaim) (bool, error) {

	reader := bufio.NewReader(opt.Options.IOStreams.In)
	for {
		var b bytes.Buffer

		var groupResource string
		if p.GroupResource.Group != "" {
			groupResource = fmt.Sprintf("%s objects (apiVersion: \"%s/%s\")", p.GroupResource.Resource, p.GroupResource.Group, p.Version)
		} else {
			groupResource = fmt.Sprintf("%s objects (apiVersion: \"%s\")", p.GroupResource.Resource, p.Version)
		}
		if p.Selector.Owner == kubebindv1alpha1.Provider {
			if p.Selector.Name != "" {
				if p.Create.Donate {
					fmt.Fprintf(&b, "Create and modify the following user-owned %s on your cluster referenced with:\n	name: \"%s\"\n", groupResource, p.Selector.Name)
				} else {
					fmt.Fprintf(&b, "Create and modify the following provider owned %s on your cluster referenced with:\n	name: \"%s\"\n", groupResource, p.Selector.Name)
				}
			} else {
				if p.Create.Donate {
					fmt.Fprintf(&b, "Create and modify user-owned %s on your cluster.\n", groupResource)
				} else {
					fmt.Fprintf(&b, "Create and modify provider owned %s on your cluster.\n", groupResource)
				}

			}
			if p.OnConflict != nil {
				switch {
				case p.OnConflict.ProviderOverwrites && p.OnConflict.RecreateWhenConsumerSideDeleted:
					b.WriteString("Conflicting objects will be overwritten and created objects will be recreated upon deletion.\n")
				case p.OnConflict.ProviderOverwrites:
					b.WriteString("Conflicting objects will be overwritten and created objects will not be recreated upon deletion.\n")
				case p.OnConflict.RecreateWhenConsumerSideDeleted:
					b.WriteString("Conflicting objects will not be overwritten and created objects will be recreated upon deletion.\n")
				default: //Do nothing
				}
			}
			if p.Update != nil {
				if p.Update.Fields != nil {
					b.WriteString("The following fields of the objects will still be managed by the provider:\n")
					for _, s := range p.Update.Fields {
						fmt.Fprintf(&b, "	\"%s\"\n", s)
					}
				}
				if p.Update.Preserving != nil {
					b.WriteString("The following fields of the objects will still be managed by the user:\n")
					for _, s := range p.Update.Preserving {
						fmt.Fprintf(&b, "	\"%s\"\n", s)
					}
				}
				if p.Update.AlwaysRecreate {
					b.WriteString("Modification of said objects will by handled by deletion and recreation of said objects.\n")
				}
			}
		}
		if p.Selector.Owner == kubebindv1alpha1.Consumer {
			if p.Selector.Name != "" {
				if p.Adopt {
					fmt.Fprintf(&b, "The provider wants to become owner of the following user-created %s referenced with:\n	name: \"%s\"\n", groupResource, p.Selector.Name)
					b.WriteString("The provider will be able to access, modify and delete said objects on your cluster.\n")
				} else {
					fmt.Fprintf(&b, "The provider wants read access to the following user-created %s referenced with:\n	name: \"%s\"\n", groupResource, p.Selector.Name)
					b.WriteString("The provider will not be able to modify or delete said objects.\n")
				}
			} else {
				if p.Adopt {
					fmt.Fprintf(&b, "The provider wants to become owner of all user-created %s.\n", groupResource)
					b.WriteString("The provider will be able to access, modify and delete said objects on your cluster.\n")
				} else {
					fmt.Fprintf(&b, "The provider wants read access to all user-created %s.\n", groupResource)
					b.WriteString("The provider will not be able to modify or delete said objects.\n")
				}
			}

			if p.Update != nil {
				if (!p.Adopt && p.Update.Fields != nil) || (p.Adopt && p.Update.Preserving != nil) {
					b.WriteString("The following fields of the objects will still be managed by the user:\n")
				}
				if (!p.Adopt && p.Update.Preserving != nil) || (p.Adopt && p.Update.Fields != nil) {
					b.WriteString("The following fields of the objects will still be managed by the provider:\n")
				}
				for _, s := range append(p.Update.Fields, p.Update.Preserving...) {
					fmt.Fprintf(&b, "	\"%s\"\n", s)
				}
				if p.Update.AlwaysRecreate {
					b.WriteString("Modification of said objects will by handled by deletion and recreation of said objects.\n")
				}
			}
		}

		// Output the prior built string
		fmt.Fprint(opt.Options.IOStreams.Out, b.String())

		if p.Required {
			fmt.Fprint(opt.Options.IOStreams.Out, "Accepting this Permission is required in order to proceed.\n")
		} else {
			fmt.Fprint(opt.Options.IOStreams.Out, "Accepting this Permission is optional.\n")
		}
		fmt.Fprint(opt.Options.IOStreams.Out, "Do you accept this Permission? [No,Yes]\n")

		response, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true, nil
		} else if response == "n" || response == "no" {
			return false, nil
		}
	}
}
