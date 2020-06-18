// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Note: this is a scaled down version of the actual CRD for Gateway
//  in order to generate some informer and lister code

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Port struct {
	Name     string `json:"name,omitempty"`
	Number   uint32 `json:"number,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

type Server struct {
	Hosts []string `json:"hosts,omitempty"`
	Port  Port     `json:"port,omitempty"`
}

type GatewaySpec struct {
	Servers  []Server          `json:"servers,omitempty"`
	Selector map[string]string `json:"selector"`
}

type GatewayStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +genclient
// +genclient:noStatus
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewaySpec   `json:"spec"`
	Status GatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
