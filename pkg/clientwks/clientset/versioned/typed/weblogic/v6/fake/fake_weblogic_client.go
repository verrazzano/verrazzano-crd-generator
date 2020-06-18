// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v6 "github.com/verrazzano/verrazzano-crd-generator/pkg/clientwks/clientset/versioned/typed/weblogic/v6"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeWeblogicV6 struct {
	*testing.Fake
}

func (c *FakeWeblogicV6) Domains(namespace string) v6.DomainInterface {
	return &FakeDomains{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeWeblogicV6) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
