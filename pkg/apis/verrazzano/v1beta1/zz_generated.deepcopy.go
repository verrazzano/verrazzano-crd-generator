// +build !ignore_autogenerated

// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingComponent) DeepCopyInto(out *BindingComponent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingComponent.
func (in *BindingComponent) DeepCopy() *BindingComponent {
	if in == nil {
		return nil
	}
	out := new(BindingComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNamespace) DeepCopyInto(out *KubernetesNamespace) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]BindingComponent, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNamespace.
func (in *KubernetesNamespace) DeepCopy() *KubernetesNamespace {
	if in == nil {
		return nil
	}
	out := new(KubernetesNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchRequest) DeepCopyInto(out *MatchRequest) {
	*out = *in
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchRequest.
func (in *MatchRequest) DeepCopy() *MatchRequest {
	if in == nil {
		return nil
	}
	out := new(MatchRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoBinding) DeepCopyInto(out *VerrazzanoBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoBinding.
func (in *VerrazzanoBinding) DeepCopy() *VerrazzanoBinding {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoBindingList) DeepCopyInto(out *VerrazzanoBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VerrazzanoBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoBindingList.
func (in *VerrazzanoBindingList) DeepCopy() *VerrazzanoBindingList {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoBindingSpec) DeepCopyInto(out *VerrazzanoBindingSpec) {
	*out = *in
	if in.WeblogicBindings != nil {
		in, out := &in.WeblogicBindings, &out.WeblogicBindings
		*out = make([]VerrazzanoWeblogicBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CoherenceBindings != nil {
		in, out := &in.CoherenceBindings, &out.CoherenceBindings
		*out = make([]VerrazzanoCoherenceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HelidonBindings != nil {
		in, out := &in.HelidonBindings, &out.HelidonBindings
		*out = make([]VerrazzanoHelidonBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DatabaseBindings != nil {
		in, out := &in.DatabaseBindings, &out.DatabaseBindings
		*out = make([]VerrazzanoDatabaseBinding, len(*in))
		copy(*out, *in)
	}
	if in.IngressBindings != nil {
		in, out := &in.IngressBindings, &out.IngressBindings
		*out = make([]VerrazzanoIngressBinding, len(*in))
		copy(*out, *in)
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = make([]VerrazzanoPlacement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoBindingSpec.
func (in *VerrazzanoBindingSpec) DeepCopy() *VerrazzanoBindingSpec {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoBindingStatus) DeepCopyInto(out *VerrazzanoBindingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoBindingStatus.
func (in *VerrazzanoBindingStatus) DeepCopy() *VerrazzanoBindingStatus {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoCoherenceBinding) DeepCopyInto(out *VerrazzanoCoherenceBinding) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoCoherenceBinding.
func (in *VerrazzanoCoherenceBinding) DeepCopy() *VerrazzanoCoherenceBinding {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoCoherenceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoCoherenceCluster) DeepCopyInto(out *VerrazzanoCoherenceCluster) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]VerrazzanoConnections, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Metrics = in.Metrics
	out.Logging = in.Logging
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoCoherenceCluster.
func (in *VerrazzanoCoherenceCluster) DeepCopy() *VerrazzanoCoherenceCluster {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoCoherenceCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoCoherenceConnection) DeepCopyInto(out *VerrazzanoCoherenceConnection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoCoherenceConnection.
func (in *VerrazzanoCoherenceConnection) DeepCopy() *VerrazzanoCoherenceConnection {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoCoherenceConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoConnections) DeepCopyInto(out *VerrazzanoConnections) {
	*out = *in
	if in.Rest != nil {
		in, out := &in.Rest, &out.Rest
		*out = make([]VerrazzanoRestConnection, len(*in))
		copy(*out, *in)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]VerrazzanoIngressConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = make([]VerrazzanoDatabaseConnection, len(*in))
		copy(*out, *in)
	}
	if in.Coherence != nil {
		in, out := &in.Coherence, &out.Coherence
		*out = make([]VerrazzanoCoherenceConnection, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoConnections.
func (in *VerrazzanoConnections) DeepCopy() *VerrazzanoConnections {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoConnections)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoDatabaseBinding) DeepCopyInto(out *VerrazzanoDatabaseBinding) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoDatabaseBinding.
func (in *VerrazzanoDatabaseBinding) DeepCopy() *VerrazzanoDatabaseBinding {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoDatabaseBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoDatabaseConnection) DeepCopyInto(out *VerrazzanoDatabaseConnection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoDatabaseConnection.
func (in *VerrazzanoDatabaseConnection) DeepCopy() *VerrazzanoDatabaseConnection {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoDatabaseConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoHelidon) DeepCopyInto(out *VerrazzanoHelidon) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.FluentdEnabled != nil {
		in, out := &in.FluentdEnabled, &out.FluentdEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]VerrazzanoConnections, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Metrics = in.Metrics
	out.Logging = in.Logging
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoHelidon.
func (in *VerrazzanoHelidon) DeepCopy() *VerrazzanoHelidon {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoHelidon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoHelidonBinding) DeepCopyInto(out *VerrazzanoHelidonBinding) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoHelidonBinding.
func (in *VerrazzanoHelidonBinding) DeepCopy() *VerrazzanoHelidonBinding {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoHelidonBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoIngressBinding) DeepCopyInto(out *VerrazzanoIngressBinding) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoIngressBinding.
func (in *VerrazzanoIngressBinding) DeepCopy() *VerrazzanoIngressBinding {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoIngressBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoIngressConnection) DeepCopyInto(out *VerrazzanoIngressConnection) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]MatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoIngressConnection.
func (in *VerrazzanoIngressConnection) DeepCopy() *VerrazzanoIngressConnection {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoIngressConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoLogging) DeepCopyInto(out *VerrazzanoLogging) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoLogging.
func (in *VerrazzanoLogging) DeepCopy() *VerrazzanoLogging {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoManagedCluster) DeepCopyInto(out *VerrazzanoManagedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoManagedCluster.
func (in *VerrazzanoManagedCluster) DeepCopy() *VerrazzanoManagedCluster {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoManagedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoManagedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoManagedClusterList) DeepCopyInto(out *VerrazzanoManagedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VerrazzanoManagedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoManagedClusterList.
func (in *VerrazzanoManagedClusterList) DeepCopy() *VerrazzanoManagedClusterList {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoManagedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoManagedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoManagedClusterSpec) DeepCopyInto(out *VerrazzanoManagedClusterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoManagedClusterSpec.
func (in *VerrazzanoManagedClusterSpec) DeepCopy() *VerrazzanoManagedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoManagedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoManagedClusterStatus) DeepCopyInto(out *VerrazzanoManagedClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoManagedClusterStatus.
func (in *VerrazzanoManagedClusterStatus) DeepCopy() *VerrazzanoManagedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoManagedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoMetrics) DeepCopyInto(out *VerrazzanoMetrics) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoMetrics.
func (in *VerrazzanoMetrics) DeepCopy() *VerrazzanoMetrics {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoModel) DeepCopyInto(out *VerrazzanoModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoModel.
func (in *VerrazzanoModel) DeepCopy() *VerrazzanoModel {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoModelList) DeepCopyInto(out *VerrazzanoModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VerrazzanoModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoModelList.
func (in *VerrazzanoModelList) DeepCopy() *VerrazzanoModelList {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoModelSpec) DeepCopyInto(out *VerrazzanoModelSpec) {
	*out = *in
	if in.WeblogicDomains != nil {
		in, out := &in.WeblogicDomains, &out.WeblogicDomains
		*out = make([]VerrazzanoWebLogicDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CoherenceClusters != nil {
		in, out := &in.CoherenceClusters, &out.CoherenceClusters
		*out = make([]VerrazzanoCoherenceCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HelidonApplications != nil {
		in, out := &in.HelidonApplications, &out.HelidonApplications
		*out = make([]VerrazzanoHelidon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoModelSpec.
func (in *VerrazzanoModelSpec) DeepCopy() *VerrazzanoModelSpec {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoModelStatus) DeepCopyInto(out *VerrazzanoModelStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoModelStatus.
func (in *VerrazzanoModelStatus) DeepCopy() *VerrazzanoModelStatus {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoPlacement) DeepCopyInto(out *VerrazzanoPlacement) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]KubernetesNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoPlacement.
func (in *VerrazzanoPlacement) DeepCopy() *VerrazzanoPlacement {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoRestConnection) DeepCopyInto(out *VerrazzanoRestConnection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoRestConnection.
func (in *VerrazzanoRestConnection) DeepCopy() *VerrazzanoRestConnection {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoRestConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoWebLogicDomain) DeepCopyInto(out *VerrazzanoWebLogicDomain) {
	*out = *in
	in.DomainCRValues.DeepCopyInto(&out.DomainCRValues)
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]VerrazzanoConnections, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Metrics = in.Metrics
	out.Logging = in.Logging
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoWebLogicDomain.
func (in *VerrazzanoWebLogicDomain) DeepCopy() *VerrazzanoWebLogicDomain {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoWebLogicDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoWeblogicBinding) DeepCopyInto(out *VerrazzanoWeblogicBinding) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoWeblogicBinding.
func (in *VerrazzanoWeblogicBinding) DeepCopy() *VerrazzanoWeblogicBinding {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoWeblogicBinding)
	in.DeepCopyInto(out)
	return out
}
