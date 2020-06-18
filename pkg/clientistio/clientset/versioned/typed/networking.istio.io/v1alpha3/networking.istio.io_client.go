// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/networking.istio.io/v1alpha3"
	"github.com/verrazzano/verrazzano-crd-generator/pkg/clientistio/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type NetworkingV1alpha3Interface interface {
	RESTClient() rest.Interface
	GatewaysGetter
	ServiceEntriesGetter
	VirtualServicesGetter
}

// NetworkingV1alpha3Client is used to interact with features provided by the networking.istio.io group.
type NetworkingV1alpha3Client struct {
	restClient rest.Interface
}

func (c *NetworkingV1alpha3Client) Gateways(namespace string) GatewayInterface {
	return newGateways(c, namespace)
}

func (c *NetworkingV1alpha3Client) ServiceEntries(namespace string) ServiceEntryInterface {
	return newServiceEntries(c, namespace)
}

func (c *NetworkingV1alpha3Client) VirtualServices(namespace string) VirtualServiceInterface {
	return newVirtualServices(c, namespace)
}

// NewForConfig creates a new NetworkingV1alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*NetworkingV1alpha3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkingV1alpha3Client{client}, nil
}

// NewForConfigOrDie creates a new NetworkingV1alpha3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkingV1alpha3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NetworkingV1alpha3Client for the given RESTClient.
func New(c rest.Interface) *NetworkingV1alpha3Client {
	return &NetworkingV1alpha3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha3.SchemeGroupVersion
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
func (c *NetworkingV1alpha3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}