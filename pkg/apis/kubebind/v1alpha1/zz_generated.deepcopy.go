//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	conditionsv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceBinding) DeepCopyInto(out *APIServiceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceBinding.
func (in *APIServiceBinding) DeepCopy() *APIServiceBinding {
	if in == nil {
		return nil
	}
	out := new(APIServiceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceBindingList) DeepCopyInto(out *APIServiceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIServiceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceBindingList.
func (in *APIServiceBindingList) DeepCopy() *APIServiceBindingList {
	if in == nil {
		return nil
	}
	out := new(APIServiceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceBindingSpec) DeepCopyInto(out *APIServiceBindingSpec) {
	*out = *in
	out.KubeconfigSecretRef = in.KubeconfigSecretRef
	if in.PermissionClaims != nil {
		in, out := &in.PermissionClaims, &out.PermissionClaims
		*out = make([]AcceptablePermissionClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceBindingSpec.
func (in *APIServiceBindingSpec) DeepCopy() *APIServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(APIServiceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceBindingStatus) DeepCopyInto(out *APIServiceBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceBindingStatus.
func (in *APIServiceBindingStatus) DeepCopy() *APIServiceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(APIServiceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExport) DeepCopyInto(out *APIServiceExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExport.
func (in *APIServiceExport) DeepCopy() *APIServiceExport {
	if in == nil {
		return nil
	}
	out := new(APIServiceExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportCRDSpec) DeepCopyInto(out *APIServiceExportCRDSpec) {
	*out = *in
	in.Names.DeepCopyInto(&out.Names)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]APIServiceExportVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportCRDSpec.
func (in *APIServiceExportCRDSpec) DeepCopy() *APIServiceExportCRDSpec {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportCRDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportList) DeepCopyInto(out *APIServiceExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIServiceExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportList.
func (in *APIServiceExportList) DeepCopy() *APIServiceExportList {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportRequest) DeepCopyInto(out *APIServiceExportRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportRequest.
func (in *APIServiceExportRequest) DeepCopy() *APIServiceExportRequest {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceExportRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportRequestList) DeepCopyInto(out *APIServiceExportRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIServiceExportRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportRequestList.
func (in *APIServiceExportRequestList) DeepCopy() *APIServiceExportRequestList {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceExportRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportRequestResource) DeepCopyInto(out *APIServiceExportRequestResource) {
	*out = *in
	out.GroupResource = in.GroupResource
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PermissionClaims != nil {
		in, out := &in.PermissionClaims, &out.PermissionClaims
		*out = make([]PermissionClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportRequestResource.
func (in *APIServiceExportRequestResource) DeepCopy() *APIServiceExportRequestResource {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportRequestResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportRequestResponse) DeepCopyInto(out *APIServiceExportRequestResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportRequestResponse.
func (in *APIServiceExportRequestResponse) DeepCopy() *APIServiceExportRequestResponse {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportRequestResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceExportRequestResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportRequestSpec) DeepCopyInto(out *APIServiceExportRequestSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]APIServiceExportRequestResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportRequestSpec.
func (in *APIServiceExportRequestSpec) DeepCopy() *APIServiceExportRequestSpec {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportRequestStatus) DeepCopyInto(out *APIServiceExportRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportRequestStatus.
func (in *APIServiceExportRequestStatus) DeepCopy() *APIServiceExportRequestStatus {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportSchema) DeepCopyInto(out *APIServiceExportSchema) {
	*out = *in
	in.OpenAPIV3Schema.DeepCopyInto(&out.OpenAPIV3Schema)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportSchema.
func (in *APIServiceExportSchema) DeepCopy() *APIServiceExportSchema {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportSpec) DeepCopyInto(out *APIServiceExportSpec) {
	*out = *in
	in.APIServiceExportCRDSpec.DeepCopyInto(&out.APIServiceExportCRDSpec)
	if in.PermissionClaims != nil {
		in, out := &in.PermissionClaims, &out.PermissionClaims
		*out = make([]PermissionClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportSpec.
func (in *APIServiceExportSpec) DeepCopy() *APIServiceExportSpec {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportStatus) DeepCopyInto(out *APIServiceExportStatus) {
	*out = *in
	in.AcceptedNames.DeepCopyInto(&out.AcceptedNames)
	if in.StoredVersions != nil {
		in, out := &in.StoredVersions, &out.StoredVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportStatus.
func (in *APIServiceExportStatus) DeepCopy() *APIServiceExportStatus {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceExportVersion) DeepCopyInto(out *APIServiceExportVersion) {
	*out = *in
	if in.DeprecationWarning != nil {
		in, out := &in.DeprecationWarning, &out.DeprecationWarning
		*out = new(string)
		**out = **in
	}
	in.Schema.DeepCopyInto(&out.Schema)
	in.Subresources.DeepCopyInto(&out.Subresources)
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]v1.CustomResourceColumnDefinition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceExportVersion.
func (in *APIServiceExportVersion) DeepCopy() *APIServiceExportVersion {
	if in == nil {
		return nil
	}
	out := new(APIServiceExportVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceNamespace) DeepCopyInto(out *APIServiceNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceNamespace.
func (in *APIServiceNamespace) DeepCopy() *APIServiceNamespace {
	if in == nil {
		return nil
	}
	out := new(APIServiceNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceNamespaceList) DeepCopyInto(out *APIServiceNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIServiceNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceNamespaceList.
func (in *APIServiceNamespaceList) DeepCopy() *APIServiceNamespaceList {
	if in == nil {
		return nil
	}
	out := new(APIServiceNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceNamespaceSpec) DeepCopyInto(out *APIServiceNamespaceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceNamespaceSpec.
func (in *APIServiceNamespaceSpec) DeepCopy() *APIServiceNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(APIServiceNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceNamespaceStatus) DeepCopyInto(out *APIServiceNamespaceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceNamespaceStatus.
func (in *APIServiceNamespaceStatus) DeepCopy() *APIServiceNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(APIServiceNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptablePermissionClaim) DeepCopyInto(out *AcceptablePermissionClaim) {
	*out = *in
	in.PermissionClaim.DeepCopyInto(&out.PermissionClaim)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptablePermissionClaim.
func (in *AcceptablePermissionClaim) DeepCopy() *AcceptablePermissionClaim {
	if in == nil {
		return nil
	}
	out := new(AcceptablePermissionClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationMethod) DeepCopyInto(out *AuthenticationMethod) {
	*out = *in
	if in.OAuth2CodeGrant != nil {
		in, out := &in.OAuth2CodeGrant, &out.OAuth2CodeGrant
		*out = new(OAuth2CodeGrant)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationMethod.
func (in *AuthenticationMethod) DeepCopy() *AuthenticationMethod {
	if in == nil {
		return nil
	}
	out := new(AuthenticationMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingProvider) DeepCopyInto(out *BindingProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.AuthenticationMethods != nil {
		in, out := &in.AuthenticationMethods, &out.AuthenticationMethods
		*out = make([]AuthenticationMethod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingProvider.
func (in *BindingProvider) DeepCopy() *BindingProvider {
	if in == nil {
		return nil
	}
	out := new(BindingProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingResponse) DeepCopyInto(out *BindingResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Authentication.DeepCopyInto(&out.Authentication)
	if in.Kubeconfig != nil {
		in, out := &in.Kubeconfig, &out.Kubeconfig
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingResponse.
func (in *BindingResponse) DeepCopy() *BindingResponse {
	if in == nil {
		return nil
	}
	out := new(BindingResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingResponseAuthentication) DeepCopyInto(out *BindingResponseAuthentication) {
	*out = *in
	if in.OAuth2CodeGrant != nil {
		in, out := &in.OAuth2CodeGrant, &out.OAuth2CodeGrant
		*out = new(BindingResponseAuthenticationOAuth2CodeGrant)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingResponseAuthentication.
func (in *BindingResponseAuthentication) DeepCopy() *BindingResponseAuthentication {
	if in == nil {
		return nil
	}
	out := new(BindingResponseAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingResponseAuthenticationOAuth2CodeGrant) DeepCopyInto(out *BindingResponseAuthenticationOAuth2CodeGrant) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingResponseAuthenticationOAuth2CodeGrant.
func (in *BindingResponseAuthenticationOAuth2CodeGrant) DeepCopy() *BindingResponseAuthenticationOAuth2CodeGrant {
	if in == nil {
		return nil
	}
	out := new(BindingResponseAuthenticationOAuth2CodeGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBinding) DeepCopyInto(out *ClusterBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBinding.
func (in *ClusterBinding) DeepCopy() *ClusterBinding {
	if in == nil {
		return nil
	}
	out := new(ClusterBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBindingList) DeepCopyInto(out *ClusterBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBindingList.
func (in *ClusterBindingList) DeepCopy() *ClusterBindingList {
	if in == nil {
		return nil
	}
	out := new(ClusterBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBindingSpec) DeepCopyInto(out *ClusterBindingSpec) {
	*out = *in
	out.KubeconfigSecretRef = in.KubeconfigSecretRef
	in.ServiceProviderSpec.DeepCopyInto(&out.ServiceProviderSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBindingSpec.
func (in *ClusterBindingSpec) DeepCopy() *ClusterBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBindingStatus) DeepCopyInto(out *ClusterBindingStatus) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	out.HeartbeatInterval = in.HeartbeatInterval
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBindingStatus.
func (in *ClusterBindingStatus) DeepCopy() *ClusterBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretKeyRef) DeepCopyInto(out *ClusterSecretKeyRef) {
	*out = *in
	out.LocalSecretKeyRef = in.LocalSecretKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretKeyRef.
func (in *ClusterSecretKeyRef) DeepCopy() *ClusterSecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionClaimCreateOptions) DeepCopyInto(out *PermissionClaimCreateOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionClaimCreateOptions.
func (in *PermissionClaimCreateOptions) DeepCopy() *PermissionClaimCreateOptions {
	if in == nil {
		return nil
	}
	out := new(PermissionClaimCreateOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSecretKeyRef) DeepCopyInto(out *LocalSecretKeyRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSecretKeyRef.
func (in *LocalSecretKeyRef) DeepCopy() *LocalSecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(LocalSecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Matcher) DeepCopyInto(out *Matcher) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Matcher.
func (in *Matcher) DeepCopy() *Matcher {
	if in == nil {
		return nil
	}
	out := new(Matcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameObjectMeta) DeepCopyInto(out *NameObjectMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameObjectMeta.
func (in *NameObjectMeta) DeepCopy() *NameObjectMeta {
	if in == nil {
		return nil
	}
	out := new(NameObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth2CodeGrant) DeepCopyInto(out *OAuth2CodeGrant) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth2CodeGrant.
func (in *OAuth2CodeGrant) DeepCopy() *OAuth2CodeGrant {
	if in == nil {
		return nil
	}
	out := new(OAuth2CodeGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionClaimOnConflictOptions) DeepCopyInto(out *PermissionClaimOnConflictOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionClaimOnConflictOptions.
func (in *PermissionClaimOnConflictOptions) DeepCopy() *PermissionClaimOnConflictOptions {
	if in == nil {
		return nil
	}
	out := new(PermissionClaimOnConflictOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionClaim) DeepCopyInto(out *PermissionClaim) {
	*out = *in
	out.GroupResource = in.GroupResource
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(ResourceSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(PermissionClaimReadOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(PermissionClaimCreateOptions)
		**out = **in
	}
	if in.OnConflict != nil {
		in, out := &in.OnConflict, &out.OnConflict
		*out = new(PermissionClaimOnConflictOptions)
		**out = **in
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		*out = new(PermissionClaimUpdateOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionClaim.
func (in *PermissionClaim) DeepCopy() *PermissionClaim {
	if in == nil {
		return nil
	}
	out := new(PermissionClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionClaimReadOptions) DeepCopyInto(out *PermissionClaimReadOptions) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.LabelsOnProviderOwnedObjects != nil {
		in, out := &in.LabelsOnProviderOwnedObjects, &out.LabelsOnProviderOwnedObjects
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.OverrideAnnotations != nil {
		in, out := &in.OverrideAnnotations, &out.OverrideAnnotations
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionClaimReadOptions.
func (in *PermissionClaimReadOptions) DeepCopy() *PermissionClaimReadOptions {
	if in == nil {
		return nil
	}
	out := new(PermissionClaimReadOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.FieldSelectors != nil {
		in, out := &in.FieldSelectors, &out.FieldSelectors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionClaimUpdateOptions) DeepCopyInto(out *PermissionClaimUpdateOptions) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Preserving != nil {
		in, out := &in.Preserving, &out.Preserving
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.LabelsOnConsumerOwnedObjects != nil {
		in, out := &in.LabelsOnConsumerOwnedObjects, &out.LabelsOnConsumerOwnedObjects
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.AnnotationsOnConsumerOwnedObjects != nil {
		in, out := &in.AnnotationsOnConsumerOwnedObjects, &out.AnnotationsOnConsumerOwnedObjects
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionClaimUpdateOptions.
func (in *PermissionClaimUpdateOptions) DeepCopy() *PermissionClaimUpdateOptions {
	if in == nil {
		return nil
	}
	out := new(PermissionClaimUpdateOptions)
	in.DeepCopyInto(out)
	return out
}
