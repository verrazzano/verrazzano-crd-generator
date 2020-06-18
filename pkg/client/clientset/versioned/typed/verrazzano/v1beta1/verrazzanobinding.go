// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/verrazzano/v1beta1"
	scheme "github.com/verrazzano/verrazzano-crd-generator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VerrazzanoBindingsGetter has a method to return a VerrazzanoBindingInterface.
// A group's client should implement this interface.
type VerrazzanoBindingsGetter interface {
	VerrazzanoBindings(namespace string) VerrazzanoBindingInterface
}

// VerrazzanoBindingInterface has methods to work with VerrazzanoBinding resources.
type VerrazzanoBindingInterface interface {
	Create(*v1beta1.VerrazzanoBinding) (*v1beta1.VerrazzanoBinding, error)
	Update(*v1beta1.VerrazzanoBinding) (*v1beta1.VerrazzanoBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.VerrazzanoBinding, error)
	List(opts v1.ListOptions) (*v1beta1.VerrazzanoBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VerrazzanoBinding, err error)
	VerrazzanoBindingExpansion
}

// verrazzanoBindings implements VerrazzanoBindingInterface
type verrazzanoBindings struct {
	client rest.Interface
	ns     string
}

// newVerrazzanoBindings returns a VerrazzanoBindings
func newVerrazzanoBindings(c *VerrazzanoV1beta1Client, namespace string) *verrazzanoBindings {
	return &verrazzanoBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the verrazzanoBinding, and returns the corresponding verrazzanoBinding object, and an error if there is any.
func (c *verrazzanoBindings) Get(name string, options v1.GetOptions) (result *v1beta1.VerrazzanoBinding, err error) {
	result = &v1beta1.VerrazzanoBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VerrazzanoBindings that match those selectors.
func (c *verrazzanoBindings) List(opts v1.ListOptions) (result *v1beta1.VerrazzanoBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VerrazzanoBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested verrazzanoBindings.
func (c *verrazzanoBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a verrazzanoBinding and creates it.  Returns the server's representation of the verrazzanoBinding, and an error, if there is any.
func (c *verrazzanoBindings) Create(verrazzanoBinding *v1beta1.VerrazzanoBinding) (result *v1beta1.VerrazzanoBinding, err error) {
	result = &v1beta1.VerrazzanoBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		Body(verrazzanoBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a verrazzanoBinding and updates it. Returns the server's representation of the verrazzanoBinding, and an error, if there is any.
func (c *verrazzanoBindings) Update(verrazzanoBinding *v1beta1.VerrazzanoBinding) (result *v1beta1.VerrazzanoBinding, err error) {
	result = &v1beta1.VerrazzanoBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		Name(verrazzanoBinding.Name).
		Body(verrazzanoBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the verrazzanoBinding and deletes it. Returns an error if one occurs.
func (c *verrazzanoBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *verrazzanoBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("verrazzanobindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched verrazzanoBinding.
func (c *verrazzanoBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VerrazzanoBinding, err error) {
	result = &v1beta1.VerrazzanoBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("verrazzanobindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
