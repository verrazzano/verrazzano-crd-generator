// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v6

import (
	v6 "github.com/oracle/verrazzano-crd-generator/pkg/apis/weblogic/v6"
	"github.com/oracle/verrazzano-crd-generator/pkg/clientwks/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type WeblogicV6Interface interface {
	RESTClient() rest.Interface
	DomainsGetter
}

// WeblogicV6Client is used to interact with features provided by the weblogic.oracle group.
type WeblogicV6Client struct {
	restClient rest.Interface
}

func (c *WeblogicV6Client) Domains(namespace string) DomainInterface {
	return newDomains(c, namespace)
}

// NewForConfig creates a new WeblogicV6Client for the given config.
func NewForConfig(c *rest.Config) (*WeblogicV6Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &WeblogicV6Client{client}, nil
}

// NewForConfigOrDie creates a new WeblogicV6Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WeblogicV6Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new WeblogicV6Client for the given RESTClient.
func New(c rest.Interface) *WeblogicV6Client {
	return &WeblogicV6Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v6.SchemeGroupVersion
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
func (c *WeblogicV6Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
