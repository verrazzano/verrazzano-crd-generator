// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package apis

import (
	v2 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/weblogic/v6"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v2.SchemeBuilder.AddToScheme)
}
