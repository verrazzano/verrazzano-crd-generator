// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v1beta1

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSockShopBinding(t *testing.T) {
	binding, err := ReadBinding("testdata/sockshop-binding.yaml")
	if err != nil {
		t.Fatal(err)
	}
	ingressBindings := binding.Spec.IngressBindings
	assert.Equal(t, 2, len(ingressBindings), "Expected 2 IngressBinding")
	assert.Equal(t, "sockshop-frontend-ingress", ingressBindings[0].Name, "Expected heliodon ingress name")
	assert.Equal(t, "wl-ingress", ingressBindings[1].Name, "Expected wls ingress name")
}

// ReadBinding reads/unmarshalls VerrazzanoBinding yaml file into a VerrazzanoBinding
func ReadBinding(path string) (*VerrazzanoBinding, error) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	var vbnd VerrazzanoBinding
	err = yaml.Unmarshal(yamlFile, &vbnd)
	return &vbnd, err
}
