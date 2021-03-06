// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v1beta1

import (
	cohv1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/coherence/v1"
	"github.com/verrazzano/verrazzano-crd-generator/pkg/apis/weblogic/v8"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VerrazzanoRestConnection defines a REST connection
// +k8s:openapi-gen=true
type VerrazzanoRestConnection struct {
	// The name of the target component
	Target string `json:"target" yaml:"target"`

	// The dns name of the target component (its k8s service)
	EnvironmentVariableForHost string `json:"environmentVariableForHost" yaml:"environmentVariableForHost"`

	// The port for the target component
	EnvironmentVariableForPort string `json:"environmentVariableForPort" yaml:"environmentVariableForPort"`
}

// VerrazzanoIngressConnection defines a Ingress connection
// +k8s:openapi-gen=true
type VerrazzanoIngressConnection struct {
	// The name of the ingress
	Name string `json:"name" yaml:"name"`

	// List of http match rules
	Match []MatchRequest `json:"match,omitempty" yaml:"match,omitempty"`
}

// MatchRequest specifies a http match rule
// +k8s:openapi-gen=true
type MatchRequest struct {
	Uri map[string]string `json:"uri,omitempty" yaml:"uri,omitempty"`
}

// VerrazzanoDatabaseConnection defines a Database conneciton
// +k8s:openapi-gen=true
type VerrazzanoDatabaseConnection struct {
	// The name of the target component
	Target string `json:"target" yaml:"target"`

	// The JDBC data source name for the database
	DatasourceName string `json:"datasourceName" yaml:"datasourceName"`
}

// VerrazzanoCoherenceConnection defines a Coherence connection
// +k8s:openapi-gen=true
type VerrazzanoCoherenceConnection struct {
	// The name of the target component
	Target string `json:"target" yaml:"target"`

	// The coherence cluster services address
	Address string `json:"address" yaml:"address"`
}

// VerrazzanoLogging defines the logging configuration
// +k8s:openapi-gen=true
type VerrazzanoLogging struct {
	// The type of logging - one of: exporter, filebeat, TBD
	Type string `json:"type" yaml:"type"`

	// The index-pattern to use (in elasticsearch)
	IndexPattern string `json:"indexPattern" yaml:"indexPattern"`
}

// VerrazzanoMetrics defines the metrics configuration for a component
// +k8s:openapi-gen=true
type VerrazzanoMetrics struct {
	// The metrics endpoint (to be scraped)
	Endpoint string `json:"endpoint" yaml:"endpoint"`

	// The name of the secret containing the credentials to access the metrics endpoint
	AuthSecret string `json:"authSecret,omitempty" yaml:"authSecret,omitempty"`

	// The interval to scrape metrics
	Interval string `json:"interval,omitempty" yaml:"interval,omitempty"`
}

// VerrazzanoConnections defines the connection for a component
// +k8s:openapi-gen=true
type VerrazzanoConnections struct {
	// REST Connections
	// +x-kubernetes-list-type=set
	Rest []VerrazzanoRestConnection `json:"rest,omitempty" yaml:"rest,omitempty"`

	// Ingress Connections
	// +x-kubernetes-list-type=set
	Ingress []VerrazzanoIngressConnection `json:"ingress,omitempty" yamml:"ingress,omitempty"`

	// Database Connections
	// +x-kubernetes-list-type=set
	Database []VerrazzanoDatabaseConnection `json:"database,omitempty" yaml:"database,omitempty"`

	// Coherence Connections
	// +x-kubernetes-list-type=set
	Coherence []VerrazzanoCoherenceConnection `json:"coherence,omitempty" yaml:"coherence,omitempty"`
}

// VerrazzanoGenericComponent defines a single generic application for the model
// +k8s:openapi-gen=true
type VerrazzanoGenericComponent struct {
	// The name of the Generic Component
	Name string `json:"name" yaml:"name"`

	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	Replicas *int32 `json:"replicas,omitempty" yaml:"replicas,omitempty"`

	// Definition of Kubernetes deployment
	Deployment corev1.PodSpec `json:"deployment" yaml:"deployment"`

	// Option to configure a deployment to use Fluentd for scraping the applications log.
	// By default, Fluentd is enabled.
	FluentdEnabled *bool `json:"fluentdEnabled,omitempty" yaml:"fluentdEnabled,omitempty"`

	// Connections configuration
	// +x-kubernetes-list-type=set
	Connections []VerrazzanoConnections `json:"connections,omitempty" yaml:"connections,omitempty"`

	// Metrics configuration
	Metrics VerrazzanoMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`
}

// VerrazzanoHelidon defines a single Helidon application for the model
// +k8s:openapi-gen=true
type VerrazzanoHelidon struct {
	// The name of the Helidon application
	Name string `json:"name" yaml:"name"`

	// The docker image:tag that runs the application
	Image string `json:"image" yaml:"image"`

	// Port to be used for the service port - defaults to 8080
	Port uint `json:"port,omitempty" yaml:"port,omitempty"`

	// TargetPort to be used for the service targetPort - defaults to 8080
	TargetPort uint `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`

	// The docker secrets for pulling images
	// +x-kubernetes-list-type=set
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`

	// Option to configure a Helidon application to use Fluentd for scraping the applications log.
	// By default, Fluentd is enabled.
	FluentdEnabled *bool `json:"fluentdEnabled,omitempty" yaml:"fluentdEnabled,omitempty"`

	// Connections configuration
	// +x-kubernetes-list-type=set
	Connections []VerrazzanoConnections `json:"connections,omitempty" yaml:"connections,omitempty"`

	// Metrics configuration
	Metrics VerrazzanoMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	// Logging configuration
	Logging VerrazzanoLogging `json:"logging,omitempty" yaml:"logging,omitempty"`

	// Environment Variables
	Env []corev1.EnvVar `json:"env,omitempty" yaml:"env,omitempty"`
}

// VerrazzanoCoherenceCluster defines a single coherence cluster for the model
// +k8s:openapi-gen=true
type VerrazzanoCoherenceCluster struct {
	// The name of the coherence cluster
	Name string `json:"name" yaml:"name"`

	// The image:tag for the coherence user artifacts
	Image string `json:"image" yaml:"image"`

	// The docker secrets for pulling images
	// +x-kubernetes-list-type=set
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`

	// Coherence pof config
	PofConfig string `json:"pofConfig" yaml:"pofConfig"`

	// Coherence cache config
	CacheConfig string `json:"cacheConfig" yaml:"cacheConfig"`

	// Connections configuration
	// +x-kubernetes-list-type=set
	Connections []VerrazzanoConnections `json:"connections,omitempty" yaml:"connections,omitempty"`

	// Coherence ports (optional)
	Ports []cohv1.NamedPortSpec `json:"ports,omitempty" yaml:"ports,omitempty"`

	// Metrics configuration
	Metrics VerrazzanoMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	// Logging configuration
	Logging VerrazzanoLogging `json:"logging,omitempty" yaml:"logging,omitempty"`
}

// VerrazzanoWebLogicDomain defines a single Weblogic domain for the model
// +k8s:openapi-gen=true
type VerrazzanoWebLogicDomain struct {
	// The WebLogic domain name
	Name string `json:"name" yaml:"name"`

	// External port number for admin console
	AdminPort int `json:"adminPort,omitempty" yaml:"adminPort,omitempty"`

	// External port number for T3
	T3Port int `json:"t3Port,omitempty" yaml:"t3Port,omitempty"`

	// Option to configure a Weblogic domain to use fluentd
	FluentdEnabled bool `json:"fluentdEnabled,omitempty" yaml:"fluentdEnabled,omitempty"`

	// Domain CR values, can provide any valid Domain CR value
	DomainCRValues v8.DomainSpec `json:"domainCRValues" yaml:"domainCRValues"`

	// Connections configuration
	// +x-kubernetes-list-type=set
	Connections []VerrazzanoConnections `json:"connections,omitempty" yaml:"connections,omitempty"`

	// Metrics configuration
	Metrics VerrazzanoMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	// Logging configuration
	Logging VerrazzanoLogging `json:"logging,omitempty" yaml:"logging,omitempty"`
}

// VerrazzanoModelSpec defines the desired state of VerrazzanoModel
// +k8s:openapi-gen=true
type VerrazzanoModelSpec struct {
	// A description of the model
	Description string `json:"description" yaml:"description"`

	// The set of WebLogic domains
	// +x-kubernetes-list-type=set
	WeblogicDomains []VerrazzanoWebLogicDomain `json:"weblogicDomains,omitempty" yaml:"weblogicDomains,omitempty"`

	// The set of coherence clusters
	// +x-kubernetes-list-type=set
	CoherenceClusters []VerrazzanoCoherenceCluster `json:"coherenceClusters,omitempty" yaml:"coherenceClusters,omitempty"`

	// The set of Helidon applications
	// +x-kubernetes-list-type=set
	HelidonApplications []VerrazzanoHelidon `json:"helidonApplications,omitempty" yaml:"helidonApplications,omitempty"`

	// The set of generic components
	// +x-kubernetes-list-type=set
	GenericComponents []VerrazzanoGenericComponent `json:"genericComponents,omitempty" yaml:"genericComponents,omitempty"`
}

// VerrazzanoModelStatus defines the observed state of VerrazzanoModel
// +k8s:openapi-gen=true
type VerrazzanoModelStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VerrazzanoModel is the Schema for the Verrazzanomodels API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=vm;vms
// +genclient
// +genclient:noStatus
type VerrazzanoModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerrazzanoModelSpec   `json:"spec"`
	Status VerrazzanoModelStatus `json:"status,omitempty"`
}

// VerrazzanoModelList contains a list of VerrazzanoModel
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VerrazzanoModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerrazzanoModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerrazzanoModel{}, &VerrazzanoModelList{})
}
