// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VerrazzanoManagedClusterSpec defines the desired state of VerrazzanoManagedCluster
// +k8s:openapi-gen=true
type VerrazzanoManagedClusterSpec struct {
	// The description of the managed cluster
	Description string `json:"description" yaml:"description"`

	// The server address
	ServerAddress string `json:"serverAddress" yaml:"serverAddress"`

	// The type of managed cluster
	Type string `json:"type" yaml:"type"`

	// The secret containing the KUBECONFIG for the managed cluster
	KubeconfigSecret string `json:"kubeconfigSecret" yaml:"kubeconfigSecret"`
}

// VerrazzanoManagedClusterStatus defines the observed state of VerrazzanoManagedCluster
// +k8s:openapi-gen=true
type VerrazzanoManagedClusterStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VerrazzanoManagedCluster is the Schema for the Verrazzanomanagedclusters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=vmc;vmcs
// +genclient
// +genclient:noStatus
type VerrazzanoManagedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerrazzanoManagedClusterSpec   `json:"spec,omitempty"`
	Status VerrazzanoManagedClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VerrazzanoManagedClusterList contains a list of VerrazzanoManagedCluster
type VerrazzanoManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerrazzanoManagedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerrazzanoManagedCluster{}, &VerrazzanoManagedClusterList{})
}
