// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// AdminServer represents the operator configuration for the admin server
// +k8s:openapi-gen=true
type AdminServer struct {
	Server `json:",inline"`

	// Configures which of the admin server's WebLogic admin channels should be exposed outside
	// the Kubernetes cluster via a node port service.
	AdminService AdminService `json:"adminService"`
}
