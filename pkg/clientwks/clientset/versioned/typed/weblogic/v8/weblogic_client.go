// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v8

import (
	v8 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/weblogic/v8"
	"github.com/verrazzano/verrazzano-crd-generator/pkg/clientwks/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type WeblogicV8Interface interface {
	RESTClient() rest.Interface
	DomainsGetter
}

// WeblogicV8Client is used to interact with features provided by the weblogic.oracle group.
type WeblogicV8Client struct {
	restClient rest.Interface
}

func (c *WeblogicV8Client) Domains(namespace string) DomainInterface {
	return newDomains(c, namespace)
}

// NewForConfig creates a new WeblogicV8Client for the given config.
func NewForConfig(c *rest.Config) (*WeblogicV8Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &WeblogicV8Client{client}, nil
}

// NewForConfigOrDie creates a new WeblogicV8Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WeblogicV8Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new WeblogicV8Client for the given RESTClient.
func New(c rest.Interface) *WeblogicV8Client {
	return &WeblogicV8Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v8.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *WeblogicV8Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
