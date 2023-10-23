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

package v1alpha1

// PermissionClaim selects objects of a GVR that a service provider may
// request and that a consumer may accept and allow the service provider access to.
//
// +kubebuilder:validation:XValidation:rule="!(has(self.autoDonate) && self.autoDonate && has(self.autoAdopt) && self.autoAdopt)",message="donate and adopt are mutually exclusive"
type PermissionClaim struct {
	GroupResource `json:","`

	// version is the version of the claimed resource.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength:=1
	Version string `json:"version"`

	// selector selects which resources are being claimed.
	// If unset, all resources across all namespaces are being claimed.
	//
	// +optional
	// +kubebuilder:default:={}
	Selector *ResourceSelector `json:"selector,omitempty"`

	// required indicates whether the APIServiceBinding will work if this claim is not accepted. If a required claim is denied, the binding is aborted.
	Required bool `json:"required"`

	// read claims read access to matching objects for the provider.
	// Reading of the claimed object(s) is always claimed.
	// By default, no labels and annotations can be read by the provider.
	// Reading of labels and annotations can be claimed in addition by specifying them explicitly.
	// If labels on consumer owned objects that are set by the consumer are read, labelsOnProviderOwnedObjects and
	// annotationsOnProviderOwnedObjects can be set.
	//
	// +optional
	// +kubebuilder:default={}
	Read *ReadOptions `json:"read,omitempty"`

	// create determines whether the kube-bind konnector will sync matching objects from the
	// provider cluster down to the consumer cluster.
	// only for owner Provider
	//
	// +optional
	Create *CreateOptions `json:"create,omitempty"`

	// autoAdopt set to true means that objects created by the consumer are adopted by the provider.
	// i.e. the provider will become the owner.
	// Mutually exclusive with autoDonate.
	//
	// +optional
	AutoAdopt bool `json:"autoAdopt,omitempty"`

	// autoDonate set to true means that a newly created object by the provider is immediately owned by the consumer.
	// If false, the object stays in ownership of the provider.
	// Mutually exclusive with autoDonate.
	//
	// +optional
	AutoDonate bool `json:"autoDonate,omitempty"`

	// onConflict determines how the conflicts between objects on the consumer cluster will be resolved.
	//
	// +optional
	// +kubebuilder:default:={}
	OnConflict *OnConflictOptions `json:"onConflict,omitempty"`

	// update lists which updates to objects on the consumer cluster are claimed.
	// By default, the whole object is synced, but metadata is not.
	//
	// +optional
	Update *UpdateOptions `json:"update,omitempty"`
}

type ReadOptions struct {
	// labels is a list of claimed label key wildcard patterns
	// that are synchronized from the consumer cluster to the provider on
	// objects that are owned by the consumer.
	//
	// +optional
	Labels []Matcher `json:"labels,omitempty"`

	// labelsOnProviderOwnedObjects is a list of claimed label key wildcard
	// patterns that are synchronized from the consumer cluster
	// to the provider on objects owned by the provider.
	//
	// +optional
	LabelsOnProviderOwnedObjects []Matcher `json:"labelsOnProviderOwnedObjects,omitempty"`

	// annotations is a list of claimed annotation key wildcard patterns
	// that are synchronized from the consumer cluster to the provider on
	// objects that are owned by the consumer.
	//
	// +optional
	Annotations []Matcher `json:"annotations,omitempty"`

	// overrideAnnotations is a list of claimed annotation key wildcard
	// patterns that are synchronized from the consumer cluster
	// to the provider on objects owned by the provider.
	//
	// +optional
	OverrideAnnotations []Matcher `json:"overrideAnnotations,omitempty"`
}

type Matcher struct {
	// +optional
	Pattern string `json:"pattern,omitempty"`
}

type OnConflictOptions struct {
	// recreateWhenConsumerSideDeleted set to true (the default) means the provider will recreate the object
	// in case the object is missing on the consumer cluster, but has been synchronized before.
	//
	// If set to false, deleted provider-owned objects get deleted on the provider cluster as well.
	//
	// +kubebuilder:default:=true
	RecreateWhenConsumerSideDeleted bool `json:"recreateWhenConsumerSideDeleted"`
}

type CreateOptions struct {
	// replaceExisting means that an existing object owned by the consumer will be replaced by the provider object.
	//
	// If not true, and a conflicting consumer object exists, it is not touched.
	//
	// +optional
	ReplaceExisting bool `json:"replaceExisting,omitempty"`
}

type UpdateOptions struct {
	// fields are a list of JSON Paths describing which parts of an object the provider wants to control.
	//
	// This field is ignored if the owner in the claim selector is set to "Provider".
	//
	// +optional
	Fields []string `json:"fields,omitempty"`

	// preserving is a list of JSON Paths describing which parts of an object owned by the provider the consumer keeps controlling.
	//
	// This field is ignored if the owner in the claim selector is set to "Consumer".
	//
	// +optional
	Preserving []string `json:"preserving,omitempty"`

	// alwaysRecreate, when true will delete the old object and create new ones
	// instead of updating. Useful for immutable objects.
	//
	// This does not apply to metadata field updates.
	//
	// +optional
	AlwaysRecreate bool `json:"alwaysRecreate,omitempty"`

	// labels is a list of claimed label keys or label wildcard patterns that are synchronized from the provider to the consumer for objects owned by the provider.
	//
	// By default, no labels are synced.
	//
	// +optional
	Labels []Matcher `json:"labels,omitempty"`

	// labelsOnConsumerOwnedObjects is a list of claimed label key wildcard patterns that are synchronized from the provider to the consumer for objects owned by the consumer.
	//
	// By default, no labels are synced.
	//
	// +optional
	LabelsOnConsumerOwnedObjects []Matcher `json:"labelsOnConsumerOwnedObjects,omitempty"`

	// annotations is a list of claimed annotation keys or annotation wildcard patterns that are synchronized from the provider to the consumer for objects owned by the provider.
	//
	// By default, no annotations are synced.
	//
	// +optional
	Annotations []Matcher `json:"annotations,omitempty"`

	// annotationsOnConsumerOwnedObjects is a list of claimed annotation key wildcard patterns that are synchronized from the provider to the consumer for objects owned by the consumer.
	//
	// By default, no annotations are synced.
	//
	// +optional
	AnnotationsOnConsumerOwnedObjects []Matcher `json:"annotationsOnConsumerOwnedObjects,omitempty"`
}

type ResourceSelector struct {
	// names is a list of specific resource names to select.
	// Names matches the metadata.name field of the underlying object.
	// An entry of "*" anywhere in the list means all object names of the group/resource within the "namespaces" field are claimed.
	// Wildcard entries other than "*" and regular expressions are currently unsupported.
	// If a resources name matches any value in names, the resource name is considered matching.
	//
	// // +kubebuilder:validation:XValidation:rule="self.all(n, n.matches('^[A-z-]+|[*]$'))",message="only names or * are allowed"
	// +kubebuilder:default:={"*"}
	// +optional
	Names []string `json:"names,omitempty"`

	// namespaces represents namespaces where an object of the given group/resource may be managed.
	// Namespaces matches against the metadata.namespace field. A value of "*" matches namespaced objects across all namespaces.
	// If a resources namespace matches any value in namespaces, the resource namespace is considered matching.
	// If the claim is for a cluster-scoped resource, namespaces has to explicitly be set to an empty array to prevent defaulting to "*".
	// If the "names" field is unset, all objects of the group/resource within the listed namespaces (or cluster) will be claimed.
	//
	// // +kubebuilder:validation:XValidation:rule="self.all(n, n.matches('^[A-z-]+|[*]$'))",message="only names or * are allowed"
	// +kubebuilder:default:={"*"}
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`

	// labelSelectors is a list of label selectors matching selected resources. label selectors follow the same rules as kubernetes label selectors,
	// see https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/.
	LabelSelectors []map[string]string `json:"labelSelectors,omitempty"`

	// fieldSelectors is a list of field selectors matching selected resources,
	// see https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/.
	FieldSelectors []string `json:"fieldSelectors,omitempty"`

	// owner matches the resource's owner. If an owner selector is set, resources owned by other owners will not be claimed.
	// Resources without a present owner will be considered, if configured owner could be the owner of the object.
	// For example, if the consumer creates a resource that is claimed by the provider for reading. In this case the resource
	// will be marked as owned by the consumer, and handled as such in further reconciliations.
	// An unset owner selector means objects from both sides are considered.
	//
	// +kubebuilder:validation:Enum=Provider;Consumer
	// +optional
	Owner Owner `json:"owner,omitempty"`
}

type Owner string

const (
	// provider means that the owner of the resource is the Provider.
	Provider Owner = "Provider"

	// consumer means that the owner of the resource is the Consumer.
	Consumer Owner = "Consumer"
)
