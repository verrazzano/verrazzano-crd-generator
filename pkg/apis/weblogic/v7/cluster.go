// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// +k8s:openapi-gen=true
type Cluster struct {
	BaseConfiguration `json:",inline"`

	// The name of this cluster
	ClusterName string `json:"clusterName"`

	// The number of managed servers to run in this cluster
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a default value.
	Replicas int `json:"replicas,omitempty"`

	// The strategy for deciding whether to start a server.
	// Legal values are NEVER, or IF_NEEDED.
	ServerStartPolicy string `json:"serverStartPolicy,omitempty"`

	// The maximum number of cluster members that can be temporarily unavailable. Defaults to 1.
	MaxUnavailable int64 `json:"maxUnavailable,omitempty"`

	// Customization affecting ClusterIP Kubernetes services for the WebLogic cluster.
	ClusterService KubernetesResource `json:"clusterService,omitempty"`
}
