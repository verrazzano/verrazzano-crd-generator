// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// DomainSpec defines the desired state of Domain
// +k8s:openapi-gen=true
type DomainSpec struct {
	BaseConfiguration `json:",inline"`

	// The name of the WebLogic domain
	DomainUID string `json:"domainUID,omitempty"`

	// The folder for the WebLogic Domain
	DomainHome string `json:"domainHome,omitempty"`

	// The strategy for deciding whether to start a server.  Legal values are ADMIN_ONLY, NEVER, or IF_NEEDED.
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a default value.
	ServerStartPolicy string `json:"serverStartPolicy,omitempty"`

	// The name of a pre-created Kubernetes secret, in the domain's namepace, that holds the username and password needed to boot WebLogic Server
	WebLogicCredentialsSecret WebLogicSecret `json:"webLogicCredentialsSecret"`

	// The in-pod name of the directory in which to store the domain, node manager, server logs, and server *.out files
	LogHome string `json:"logHome"`

	// Specified whether the log home folder is enabled
	LogHomeEnabled bool `json:"logHomeEnabled,omitempty"`

	// If true (the default), the server .out file will be included in the pod's stdout.
	IncludeServerOutInPodLog bool `json:"includeServerOutInPodLog,omitempty"`

	// The WebLogic Docker image; required when domainHomeInImage is true
	Image string `json:"image"`

	// The image pull policy for the WebLogic Docker image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// A list of image pull secrets for the WebLogic Docker image
	// +listType
	ImagePullSecrets []WebLogicSecret `json:"imagePullSecrets"`

	// The number of managed servers to run in any cluster that does not specify a replica count.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a default value.
	Replicas *int32 `json:"replicas,omitempty"`

	// True if this domain's home is defined in the docker image for the domain
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a default value.
	DomainHomeInImage bool `json:"domainHomeInImage,omitempty"`

	// The name of the config map for optional WebLogic configuration overrides
	ConfigOverrides string `json:"configOverrides,omitempty"`

	// A list of names of the secrets for optional WebLogic configuration overrides
	// +listType
	ConfigOverrideSecrets []string `json:"configOverrideSecrets,omitempty"`

	// Configuration for the admin server
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a default value.
	AdminServer AdminServer `json:"adminServer,omitempty"`

	// Configuration for individual Managed Servers
	// +listType
	ManagedServers []ManagedServer `json:"managedServers,omitempty"`

	// Configuration for the clusters
	// +listType
	Clusters []Cluster `json:"clusters"`

	// Configurations
	Configuration Configuration `json:"configuration,omitempty"`
}
