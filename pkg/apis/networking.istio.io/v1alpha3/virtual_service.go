// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Note: this is a scaled down version of the actual CRD for VirtualService
//  in order to generate some informer and lister code

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PortSelector struct {
	Number int `json:"number"`
}

type Destination struct {
	Host string       `json:"host,omitempty"`
	Port PortSelector `json:"port,omitempty"`
}

type HTTPRouteDestination struct {
	Destination Destination `json:"destination,omitempty"`
}

type MatchRequest struct {
	Uri  map[string]string `json:"uri,omitempty"`
	Port int               `json:"port,omitempty"`
}

type HttpRoute struct {
	Match []MatchRequest         `json:"match,omitempty"`
	Route []HTTPRouteDestination `json:"route,omitempty"`
}

type VirtualServiceSpec struct {
	Gateways []string    `json:"gateways,omitempty"`
	Hosts    []string    `json:"hosts,omitempty"`
	Http     []HttpRoute `json:"http,omitempty"`
}

type VirtualServiceStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +genclient
// +genclient:noStatus
type VirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualServiceSpec   `json:"spec"`
	Status VirtualServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualService{}, &VirtualServiceList{})
}
