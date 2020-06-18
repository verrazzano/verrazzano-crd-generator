// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v6

// Inline managed server definition
type ManagedServer struct {
	Server `json:",inline"`

	// The name of the Managed Server
	ServerName string `json:"serverName"`
}
