// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// Status of WebLogic servers in this domain
// +k8s:openapi-gen=true
type ServerStatus struct {
	// WebLogic server name
	ServerName string `json:"serverName"`

	// Current state of this WebLogic server
	State string `json:"state"`

	// WebLogic cluster name, if the server is part of a cluster
	ClusterName string `json:"clusterName,omitempty"`

	// Name of node that is hosting the Pod containing this WebLogic server
	NodeName string `json:"nodeName,omitempty"`

	// Current status and health of a specific WebLogic server
	Health ServerHealth `json:"health"`
}
