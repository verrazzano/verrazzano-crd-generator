// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v6

// Experimental feature configurations
// +k8s:openapi-gen=true
type Experimental struct {
	// Istio service mesh integration configuration
	Istio Istio `json:"istio,omitempty"`
}
