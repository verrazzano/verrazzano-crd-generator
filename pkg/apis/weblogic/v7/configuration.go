// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// +k8s:openapi-gen=true
type Configuration struct {
	// Istio service mesh integration configuration
	Istio Istio `json:"istio,omitempty"`
}
