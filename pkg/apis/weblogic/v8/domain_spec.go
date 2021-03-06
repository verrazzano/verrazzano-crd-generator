// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v8

import (
	corev1 "k8s.io/api/core/v1"
)

// DomainSpec defines the desired state of Domain
// +k8s:openapi-gen=true
type DomainSpec struct {
	// Configuration for the admin server
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a
	// default value.
	AdminServer AdminServer `json:"adminServer,omitempty"`

	// Whether to allow the number of running cluster member Managed Server instances to drop below the minimum
	// dynamic cluster size configured in the WebLogic domain configuration, if this is not specified for a specific
	// cluster under the clusters field. Defaults to true.
	AllowReplicasBelowMinDynClusterSize bool `json:"allowReplicasBelowMinDynClusterSize,omitempty"`

	// Configuration for the clusters
	// +x-kubernetes-list-type=set
	Clusters []Cluster `json:"clusters,omitempty"`

	// The name of the config map for optional WebLogic configuration overrides
	ConfigOverrides string `json:"configOverrides,omitempty"`

	// A list of names of the secrets for optional WebLogic configuration overrides
	// +x-kubernetes-list-type=set
	ConfigOverrideSecrets []string `json:"configOverrideSecrets,omitempty"`

	// Configurations
	Configuration Configuration `json:"configuration,omitempty"`

	// An optional directory in a server's container for data storage of default and custom file stores. If dataHome
	// is not specified or its value is either not set or empty, then the data storage directories are determined
	// from the WebLogic domain configuration.
	DataHome string `json:"dataHome,omitempty"`

	// The folder for the WebLogic Domain
	DomainHome string `json:"domainHome,omitempty"`

	// True if this domain's home is defined in the docker image for the domain
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a
	// default value.
	DomainHomeInImage bool `json:"domainHomeInImage,omitempty"`

	// Domain home file system source type: Legal values: Image, PersistentVolume, FromModel. Image indicates that
	// the domain home file system is present in the container image specified by the image field. PersistentVolume
	// indicates that the domain home file system is located on a persistent volume. FromModel indicates that the
	// domain home file system will be created and managed by the operator based on a WDT domain model. If this
	// field is specified, it overrides the value of domainHomeInImage. If both fields are unspecified, then
	// domainHomeSourceType defaults to Image.
	DomainHomeSourceType string `json:"domainHomeSourceType,omitempty"`

	// The name of the WebLogic domain
	DomainUID string `json:"domainUID,omitempty"`

	// Specifies whether the server HTTP access log files will be written to the same directory specified in logHome.
	// Otherwise, server HTTP access log files will be written to the directory configured in the WebLogic domain
	// configuration. Defaults to true.
	HttpAccessLogInLogHome bool `json:"httpAccessLogInLogHome,omitempty"`

	// The WebLogic Docker image; required when domainHomeInImage is true
	Image string `json:"image"`

	// The image pull policy for the WebLogic Docker image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// A list of image pull secrets for the WebLogic Docker image
	// +x-kubernetes-list-type=set
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets"`

	// If true (the default), the server .out file will be included in the pod's stdout.
	IncludeServerOutInPodLog bool `json:"includeServerOutInPodLog,omitempty"`

	// Changes to this field cause the operator to repeat its introspection of the WebLogic domain configuration.
	IntrospectVersion string `json:"introspectVersion,omitempty"`

	// The in-pod name of the directory in which to store the domain, node manager, server logs, and server *.out files
	LogHome string `json:"logHome"`

	// Specified whether the log home folder is enabled
	LogHomeEnabled bool `json:"logHomeEnabled,omitempty"`

	// Configuration for individual Managed Servers
	// +x-kubernetes-list-type=set
	ManagedServers []ManagedServer `json:"managedServers,omitempty"`

	// The maximum number of cluster member Managed Server instances that the operator will start in parallel
	// for a given cluster, if maxConcurrentStartup is not specified for a specific cluster under the clusters field.
	// A value of 0 means there is no configured limit. Defaults to 0.
	MaxClusterConcurrentStartup int32 `json:"maxClusterConcurrentStartup,omitempty"`

	// The number of managed servers to run in any cluster that does not specify a replica count.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a default value.
	Replicas *int32 `json:"replicas,omitempty"`

	// If present, every time this value is updated the operator will restart
	// the required servers.
	RestartVersion string `json:"restartVersion,omitempty"`

	// Server Pod
	ServerPod ServerPod `json:"serverPod,omitempty"`

	// Customization affecting ClusterIP Kubernetes services for WebLogic Server instances.
	ServerService ServerService `json:"serverService,omitempty"`

	// The strategy for deciding whether to start a server.  Legal values are ADMIN_ONLY, NEVER, or IF_NEEDED.
	// Note: this value is required by WebLogic Operator, but is marked optional because Verrazzano can provide a
	// default value.
	ServerStartPolicy string `json:"serverStartPolicy,omitempty"`

	// The state in which the server is to be started.  Legal values are "RUNNING" or "ADMIN"
	ServerStartState string `json:"serverStartState,omitempty"`

	// The name of a pre-created Kubernetes secret, in the domain's namepace, that holds the username and password
	// needed to boot WebLogic Server
	WebLogicCredentialsSecret corev1.SecretReference `json:"webLogicCredentialsSecret"`
}
