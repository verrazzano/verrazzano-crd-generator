// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v8

import (
	"context"
	"time"

	v8 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/weblogic/v8"
	scheme "github.com/verrazzano/verrazzano-crd-generator/pkg/clientwks/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DomainsGetter has a method to return a DomainInterface.
// A group's client should implement this interface.
type DomainsGetter interface {
	Domains(namespace string) DomainInterface
}

// DomainInterface has methods to work with Domain resources.
type DomainInterface interface {
	Create(ctx context.Context, domain *v8.Domain, opts v1.CreateOptions) (*v8.Domain, error)
	Update(ctx context.Context, domain *v8.Domain, opts v1.UpdateOptions) (*v8.Domain, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v8.Domain, error)
	List(ctx context.Context, opts v1.ListOptions) (*v8.DomainList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v8.Domain, err error)
	DomainExpansion
}

// domains implements DomainInterface
type domains struct {
	client rest.Interface
	ns     string
}

// newDomains returns a Domains
func newDomains(c *WeblogicV8Client, namespace string) *domains {
	return &domains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the domain, and returns the corresponding domain object, and an error if there is any.
func (c *domains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v8.Domain, err error) {
	result = &v8.Domain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("domains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Domains that match those selectors.
func (c *domains) List(ctx context.Context, opts v1.ListOptions) (result *v8.DomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v8.DomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested domains.
func (c *domains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a domain and creates it.  Returns the server's representation of the domain, and an error, if there is any.
func (c *domains) Create(ctx context.Context, domain *v8.Domain, opts v1.CreateOptions) (result *v8.Domain, err error) {
	result = &v8.Domain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(domain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a domain and updates it. Returns the server's representation of the domain, and an error, if there is any.
func (c *domains) Update(ctx context.Context, domain *v8.Domain, opts v1.UpdateOptions) (result *v8.Domain, err error) {
	result = &v8.Domain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("domains").
		Name(domain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(domain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the domain and deletes it. Returns an error if one occurs.
func (c *domains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("domains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *domains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched domain.
func (c *domains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v8.Domain, err error) {
	result = &v8.Domain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("domains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
