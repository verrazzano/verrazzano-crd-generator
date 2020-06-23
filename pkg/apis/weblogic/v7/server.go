// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// Inline server definition
type Server struct {
	BaseConfiguration `json:",inline"`

	// The strategy for deciding whether to start a server.
	// Legal values are ALWAYS, NEVER, or IF_NEEDED.
	ServerStartPolicy string `json:"serverStartPolicy,omitempty"`
}
