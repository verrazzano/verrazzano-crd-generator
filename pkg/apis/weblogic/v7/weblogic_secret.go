// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v7

// WebLogicSecret identifies the secret that contains the WebLogic admin credentials
// +k8s:openapi-gen=true
type WebLogicSecret struct {
	Name string `json:"name"`
}
