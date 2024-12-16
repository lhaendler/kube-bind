/*
Copyright 2024 The Kube Bind Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"

	kubebindv1alpha1 "github.com/kube-bind/kube-bind/sdk/apis/kubebind/v1alpha1"
)

// APIServiceBindingLister helps list APIServiceBindings.
// All objects returned here must be treated as read-only.
type APIServiceBindingLister interface {
	// List lists all APIServiceBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubebindv1alpha1.APIServiceBinding, err error)
	// Get retrieves the APIServiceBinding from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubebindv1alpha1.APIServiceBinding, error)
	APIServiceBindingListerExpansion
}

// aPIServiceBindingLister implements the APIServiceBindingLister interface.
type aPIServiceBindingLister struct {
	listers.ResourceIndexer[*kubebindv1alpha1.APIServiceBinding]
}

// NewAPIServiceBindingLister returns a new APIServiceBindingLister.
func NewAPIServiceBindingLister(indexer cache.Indexer) APIServiceBindingLister {
	return &aPIServiceBindingLister{listers.New[*kubebindv1alpha1.APIServiceBinding](indexer, kubebindv1alpha1.Resource("apiservicebinding"))}
}