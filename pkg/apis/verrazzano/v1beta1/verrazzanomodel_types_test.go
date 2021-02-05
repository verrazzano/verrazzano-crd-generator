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

func TestSockShopModel(t *testing.T) {
	model, err := ReadModel("testdata/sockshop-model.yaml")
	if err != nil {
		t.Fatal(err)
	}
	helidons := model.Spec.HelidonApplications
	assert.Equal(t, 8, len(helidons), "Expected 8 HelidonApplications")
	appf := find("frontend", helidons)
	assert.Equal(t, uint(8088), appf.Port, "Expected frontend Port")
	assert.Equal(t, uint(8079), appf.TargetPort, "Expected frontend TargetPort")
	assert.Equal(t, 1, len(appf.Connections), "Expected 1 Connection")
	assert.Equal(t, 1, len(appf.Connections[0].Ingress), "Expected 1 Ingress")
	ingr := appf.Connections[0].Ingress[0]
	assert.Equal(t, "sockshop-frontend-ingress", ingr.Name, "Expected heliodon ingress name")
	assert.Equal(t, 4, len(ingr.Match), "Expected 8 match rules")
	assert.Equal(t, "/", ingr.Match[0].Uri["prefix"], "Expected match")
	assert.Equal(t, "/cart", ingr.Match[1].Uri["exact"], "Expected match")
	assert.Equal(t, "/cart", ingr.Match[2].Uri["prefix"], "Expected match")
	assert.Equal(t, "/catalogue", ingr.Match[3].Uri["exact"], "Expected match")

	assert.Equal(t, 1, len(model.Spec.WeblogicDomains), "Expected 1 wls domeain")
	wlsd := model.Spec.WeblogicDomains[0]
	assert.Equal(t, 2, len(wlsd.Connections), "Expected 1 wls connection")
	assert.Equal(t, 1, len(wlsd.Connections[0].Ingress), "Expected 1 wls ingress")
	wingr := wlsd.Connections[0].Ingress[0]
	assert.Equal(t, "wl-ingress", wingr.Name, "Expected wls ingress name")
}

// ReadModel reads/unmarshalls VerrazzanoModel yaml file into a VerrazzanoBinding
func ReadModel(path string) (*VerrazzanoModel, error) {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	var model VerrazzanoModel
	err = yaml.Unmarshal(yamlFile, &model)
	return &model, err
}

func find(name string, helidons []VerrazzanoHelidon) VerrazzanoHelidon {
	for _, app := range helidons {
		if app.Name == name {
			return app
		}
	}
	return VerrazzanoHelidon{}
}
