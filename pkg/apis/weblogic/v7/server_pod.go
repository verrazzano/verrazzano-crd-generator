// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

import (
	corev1 "k8s.io/api/core/v1"
)

// +k8s:openapi-gen=true
type ServerPod struct {
	// A list of environment variables to add to a server
	// +x-kubernetes-list-type=set
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Settings for the liveness probe associated with a server.
	LivenessProbe ProbeTuning `json:"livenessProbe,omitempty"`

	// Settings for the readiness probe associated with a server.
	ReadinessProbe ProbeTuning `json:"readinessProbe,omitempty"`

	// Selector which must match a node's labels for the pod to be scheduled on that node.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// If specified, the pod's scheduling constraints
	Affinity corev1.Affinity `json:"affinity,omitempty"`

	// InitContainers holds a list of initialization containers that should
	// be run before starting the main containers in this pod.
	// +x-kubernetes-list-type=set
	InitContainers []corev1.Container `json:"initContainers,omitempty"`

	// Additional containers to be included in the server pod
	// +x-kubernetes-list-type=set
	Containers []corev1.Container `json:"containers,omitempty"`

	// If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical"
	// are two special keywords which indicate the highest priorities with the former being the highest priority.
	// Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod
	// priority will be default or zero if there is no default.
	PriorityClassName string `json:"priorityClassName,omitempty"`

	// If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its
	// containers are ready AND all conditions specified in the readiness gates have status equal to "True" More
	// info: https://github.com/kubernetes/community/blob/master/keps/sig-network/0007-pod-ready%2B%2B.md"
	// +x-kubernetes-list-type=set
	ReadinessGates []corev1.PodReadinessGate `json:"readinessGates,omitempty"`

	// Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always.
	RestartPolicy string `json:"restartPolicy,omitempty"`

	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run
	// this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty,
	// the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the
	// default runtime handler.
	RuntimeClassName string `json:"runtimeClassName,omitempty"`

	// NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler
	// simply schedules this pod onto that node, assuming that it fits resource requirements.
	NodeName string `json:"nodeName,omitempty"`

	// If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be
	// dispatched by default scheduler.
	SchedulerName string `json:"schedulerName,omitempty"`

	// If specified, the pod's tolerations.
	// +x-kubernetes-list-type=set
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Name of the ServiceAccount to be used to run this pod. If it is not set, default
	// ServiceAccount will be used. The ServiceAccount has to exist at the time the pod is created.
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Memory and CPU minimum requirements and limits for the server.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Pod-level security attributes.
	PodSecurityContext corev1.PodSecurityContext `json:"podSecurityContext,omitempty"`

	Shutdown Shutdown `json:"shutdown,omitempty"`

	// Additional volumes to be created in the server pod
	// +x-kubernetes-list-type=set
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Additional volume mounts for the server pod
	// +x-kubernetes-list-type=set
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
}
