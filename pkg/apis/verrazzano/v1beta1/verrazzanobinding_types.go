// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:openapi-gen=true
// BindingComponent specifies a single component to bind to a namespace
type BindingComponent struct {
	// Name of a bound component
	Name string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
// VerrazzanoPlacement defines a single placement definition
type VerrazzanoPlacement struct {
	// The name of the placement
	Name string `json:"name" yaml:"name"`

	// Namespaces for this placement
	// +listType=set
	Namespaces []KubernetesNamespace `json:"namespaces" yaml:"namespaces"`
}

// +k8s:openapi-gen=true
// KubernetesNamespace defines a single Kubernetes namespace
type KubernetesNamespace struct {
	// Name of the namespace
	Name string `json:"name" yaml:"name"`

	// Names of components in the namespace
	// +listType=set
	Components []BindingComponent `json:"components" yaml:"components"`
}

// +k8s:openapi-gen=true
// VerrazzanoIngressBinding defines a single ingress binding
type VerrazzanoIngressBinding struct {
	// The name of the binding
	Name string `json:"name" yaml:"name"`

	// The DNS name
	DnsName string `json:"dnsName" yaml:"dnsName"`
}

// +k8s:openapi-gen=true
// VerrazzanoDatabaseBinding defines a single database binding
type VerrazzanoDatabaseBinding struct {
	// The name of the binding
	Name string `json:"name" yaml:"name"`

	// The name of the secret containing the credentials
	Credentials string `json:"credentials" yaml:"credentials"`

	// The database url connection string
	Url string `json:"url" yaml:"url"`
}

// +k8s:openapi-gen=true
// VerrazzanoWeblogicBinding defines a single Weblogic domain binding
type VerrazzanoWeblogicBinding struct {
	// The name of the binding
	Name string `json:"name" yaml:"name"`

	// Number of replicas to create
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 1.
	Replicas *int32 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}

// +k8s:openapi-gen=true
// VerrazzanoCoherenceBinding defines a single Coherence binding
type VerrazzanoCoherenceBinding struct {
	// The name of the binding
	Name string `json:"name" yaml:"name"`

	// Number of replicas to create
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 3.
	Replicas *int32 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}

// +k8s:openapi-gen=true
// VerrazzanoHelidonBinding defines a single Helidon application binding
type VerrazzanoHelidonBinding struct {
	// The name of the binding
	Name string `json:"name" yaml:"name"`

	// Number of replicas to create
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 1.
	Replicas *int32 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
}

// VerrazzanoBindingSpec defines the desired state of VerrazzanoBinding
// +k8s:openapi-gen=true
type VerrazzanoBindingSpec struct {
	// A description of the binding
	Description string `json:"description" yaml:"description"`

	// The model name to associate the bindings
	ModelName string `json:"modelName" yaml:"modelName"`

	// The set of weblogic bindings
	// +listType=set
	WeblogicBindings []VerrazzanoWeblogicBinding `json:"weblogicBindings,omitempty" yaml:"weblogicBindings,omitempty"`

	// The set of coherence bindings
	// +listType=set
	CoherenceBindings []VerrazzanoCoherenceBinding `json:"coherenceBindings,omitempty" yaml:"coherenceBindings,omitempty"`

	// The set of helidon application bindings
	// +listType=set
	HelidonBindings []VerrazzanoHelidonBinding `json:"helidonBindings,omitempty" yaml:"helidonBindings,omitempty"`

	// The set of database bindings
	// +listType=set
	DatabaseBindings []VerrazzanoDatabaseBinding `json:"databaseBindings,omitempty" yaml:"databaseBindings,omitempty"`

	// The set of ingress bindings
	// +listType=set
	IngressBindings []VerrazzanoIngressBinding `json:"ingressBindings,omitempty" yaml:"ingressBindings,omitempty"`

	// The set of Placement definitions
	// +listType=set
	Placement []VerrazzanoPlacement `json:"placement" yaml:"placement"`
}

// VerrazzanoBindingStatus defines the observed state of VerrazzanoBinding
// +k8s:openapi-gen=true
type VerrazzanoBindingStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VerrazzanoBinding is the Schema for the Verrazzanobindings API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=vb;vbs
// +genclient
// +genclient:noStatus
type VerrazzanoBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerrazzanoBindingSpec   `json:"spec"`
	Status VerrazzanoBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VerrazzanoBindingList contains a list of VerrazzanoBinding
type VerrazzanoBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerrazzanoBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerrazzanoBinding{}, &VerrazzanoBindingList{})
}
