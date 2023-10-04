/*
Copyright The Kube Bind Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kube-bind/kube-bind/contrib/example-backend/apis/examplebackend/v1alpha1"
	scheme "github.com/kube-bind/kube-bind/contrib/example-backend/client/clientset/versioned/scheme"
)

// APIServiceExportTemplatesGetter has a method to return a APIServiceExportTemplateInterface.
// A group's client should implement this interface.
type APIServiceExportTemplatesGetter interface {
	APIServiceExportTemplates(namespace string) APIServiceExportTemplateInterface
}

// APIServiceExportTemplateInterface has methods to work with APIServiceExportTemplate resources.
type APIServiceExportTemplateInterface interface {
	Create(ctx context.Context, aPIServiceExportTemplate *v1alpha1.APIServiceExportTemplate, opts v1.CreateOptions) (*v1alpha1.APIServiceExportTemplate, error)
	Update(ctx context.Context, aPIServiceExportTemplate *v1alpha1.APIServiceExportTemplate, opts v1.UpdateOptions) (*v1alpha1.APIServiceExportTemplate, error)
	UpdateStatus(ctx context.Context, aPIServiceExportTemplate *v1alpha1.APIServiceExportTemplate, opts v1.UpdateOptions) (*v1alpha1.APIServiceExportTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIServiceExportTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIServiceExportTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIServiceExportTemplate, err error)
	APIServiceExportTemplateExpansion
}

// aPIServiceExportTemplates implements APIServiceExportTemplateInterface
type aPIServiceExportTemplates struct {
	client rest.Interface
	ns     string
}

// newAPIServiceExportTemplates returns a APIServiceExportTemplates
func newAPIServiceExportTemplates(c *ExampleBackendV1alpha1Client, namespace string) *aPIServiceExportTemplates {
	return &aPIServiceExportTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aPIServiceExportTemplate, and returns the corresponding aPIServiceExportTemplate object, and an error if there is any.
func (c *aPIServiceExportTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIServiceExportTemplate, err error) {
	result = &v1alpha1.APIServiceExportTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIServiceExportTemplates that match those selectors.
func (c *aPIServiceExportTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIServiceExportTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIServiceExportTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIServiceExportTemplates.
func (c *aPIServiceExportTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIServiceExportTemplate and creates it.  Returns the server's representation of the aPIServiceExportTemplate, and an error, if there is any.
func (c *aPIServiceExportTemplates) Create(ctx context.Context, aPIServiceExportTemplate *v1alpha1.APIServiceExportTemplate, opts v1.CreateOptions) (result *v1alpha1.APIServiceExportTemplate, err error) {
	result = &v1alpha1.APIServiceExportTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIServiceExportTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIServiceExportTemplate and updates it. Returns the server's representation of the aPIServiceExportTemplate, and an error, if there is any.
func (c *aPIServiceExportTemplates) Update(ctx context.Context, aPIServiceExportTemplate *v1alpha1.APIServiceExportTemplate, opts v1.UpdateOptions) (result *v1alpha1.APIServiceExportTemplate, err error) {
	result = &v1alpha1.APIServiceExportTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		Name(aPIServiceExportTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIServiceExportTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPIServiceExportTemplates) UpdateStatus(ctx context.Context, aPIServiceExportTemplate *v1alpha1.APIServiceExportTemplate, opts v1.UpdateOptions) (result *v1alpha1.APIServiceExportTemplate, err error) {
	result = &v1alpha1.APIServiceExportTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		Name(aPIServiceExportTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIServiceExportTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIServiceExportTemplate and deletes it. Returns an error if one occurs.
func (c *aPIServiceExportTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIServiceExportTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIServiceExportTemplate.
func (c *aPIServiceExportTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIServiceExportTemplate, err error) {
	result = &v1alpha1.APIServiceExportTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apiserviceexporttemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}