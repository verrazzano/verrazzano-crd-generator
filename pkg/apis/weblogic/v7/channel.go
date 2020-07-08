// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// +k8s:openapi-gen=true
type Channel struct {
	// Name of channel
	ChannelName string `json:"channelName"`

	// Specifies the port number used to access the WebLogic channel outside of the Kubernetes cluster
	NodePort int `json:"nodePort,omitempty"`
}
