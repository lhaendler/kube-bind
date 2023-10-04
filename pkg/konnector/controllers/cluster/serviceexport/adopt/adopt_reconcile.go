/*
Copyright 2023 The Kube Bind Authors.

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

package adopt

import (
	"context"

	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog/v2"

	kubebindv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

type reconciler struct {
	getServiceNamespace func(upstreamNamespace string) (*kubebindv1alpha1.APIServiceNamespace, error)

	getConsumerObject    func(ns, name string) (*unstructured.Unstructured, error)
	updateConsumerObject func(ctx context.Context, obj *unstructured.Unstructured) (*unstructured.Unstructured, error)
	createConsumerObject func(ctx context.Context, obj *unstructured.Unstructured) (*unstructured.Unstructured, error)

	updateProviderObject func(ctx context.Context, obj *unstructured.Unstructured) (*unstructured.Unstructured, error)
}

// reconcile copies new object to the consumer cluster.
func (r *reconciler) reconcile(ctx context.Context, obj *unstructured.Unstructured) error {
	logger := klog.FromContext(ctx)

	ns := obj.GetNamespace()
	if ns != "" {
		sn, err := r.getServiceNamespace(ns)
		if err != nil && !errors.IsNotFound(err) {
			return err
		} else if errors.IsNotFound(err) {
			runtime.HandleError(err)
			return err // hoping the APIServiceNamespace will be created soon. Otherwise, this item goes into backoff.
		}
		if sn.Status.Namespace == "" {
			runtime.HandleError(err)
			return err // hoping the status is set soon.
		}

		logger = logger.WithValues("providerNamespace", sn.Status.Namespace)
		ctx = klog.NewContext(ctx, logger)

		// continue with downstream namespace
		ns = sn.Name
	}

	consumerObj, err := r.getConsumerObject(ns, obj.GetName())
	if err != nil && !errors.IsNotFound(err) {
		logger.Info("failed to get downstream object", "error", err, "downstreamNamespace", ns, "downstreamName", obj.GetName())
		return err
	}

	if errors.IsNotFound(err) {
		if v, ok := obj.GetAnnotations()["kube-bind.io/bound"]; !ok && v != "true" {
			logger.Info("adopting object", "name", obj.GetName(), "consumer namespace", ns)

			//TODO does this overwrite important annotations?
			candidate := obj.DeepCopy()
			candidate.SetResourceVersion("")
			candidate.SetNamespace(ns)
			candidate.SetAnnotations(map[string]string{"kube-bind.io/bound": "true"})
			_, err := r.createConsumerObject(ctx, candidate)
			return err
		}
	}

	// Set annotation on downstream object
	cObj := consumerObj.DeepCopy()
	ann := cObj.GetAnnotations()
	if ann == nil {
		ann = map[string]string{}
	}
	ann["kube-bind.io/bound"] = "true"
	cObj.SetAnnotations(ann)

	if !equality.Semantic.DeepEqual(cObj.GetAnnotations(), consumerObj.GetAnnotations()) {
		logger.Info("adding bind annotation to consumer object", "name", cObj.GetName(), "ns", cObj.GetNamespace())
		_, err = r.updateConsumerObject(ctx, cObj)
	}

	pObj := obj.DeepCopy()
	ann = pObj.GetAnnotations()
	if ann == nil {
		ann = map[string]string{}
	}
	ann["kube-bind.io/bound"] = "true"
	pObj.SetAnnotations(ann)

	if !equality.Semantic.DeepEqual(pObj.GetAnnotations(), consumerObj.GetAnnotations()) {
		logger.Info("adding bind annotation to provider object", "name", cObj.GetName(), "ns", cObj.GetNamespace())
		_, err = r.updateProviderObject(ctx, pObj)
	}

	return err
}
