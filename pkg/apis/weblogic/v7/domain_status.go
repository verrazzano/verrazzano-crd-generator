// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// DomainStatus defines the observed state of Domain
// +k8s:openapi-gen=true
type DomainStatus struct {
	// Current service state of domain
	// +listType
	Conditions []DomainCondition `json:"conditions,omitempty"`

	// A human readable message indicating details about why the domain is in this condition
	Message string `json:"message,omitempty"`

	// A brief CamelCase message indicating details about why the domain is in this state
	Reason string `json:"reason,omitempty"`

	// Status of WebLogic servers in this domain
	// +listType
	Servers []ServerStatus `json:"servers,omitempty"`

	// RFC 3339 date and time at which the operator started the domain. This will be when
	// the operator begins processing and will precede when the various servers
	// or clusters are available.
	StartTime string `json:"startTime,omitempty"`

	// The number of running managed servers in the WebLogic cluster if there is
	// only one cluster in the domain and where the cluster does not explicitly
	// configure its replicas in a cluster specification.
	Replicas int `json:"replicas,omitempty"`
}
