// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/oracle/verrazzano-crd-generator/pkg/apis/verrazzano/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVerrazzanoModels implements VerrazzanoModelInterface
type FakeVerrazzanoModels struct {
	Fake *FakeVerrazzanoV1beta1
	ns   string
}

var verrazzanomodelsResource = schema.GroupVersionResource{Group: "verrazzano.io", Version: "v1beta1", Resource: "verrazzanomodels"}

var verrazzanomodelsKind = schema.GroupVersionKind{Group: "verrazzano.io", Version: "v1beta1", Kind: "VerrazzanoModel"}

// Get takes name of the verrazzanoModel, and returns the corresponding verrazzanoModel object, and an error if there is any.
func (c *FakeVerrazzanoModels) Get(name string, options v1.GetOptions) (result *v1beta1.VerrazzanoModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(verrazzanomodelsResource, c.ns, name), &v1beta1.VerrazzanoModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VerrazzanoModel), err
}

// List takes label and field selectors, and returns the list of VerrazzanoModels that match those selectors.
func (c *FakeVerrazzanoModels) List(opts v1.ListOptions) (result *v1beta1.VerrazzanoModelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(verrazzanomodelsResource, verrazzanomodelsKind, c.ns, opts), &v1beta1.VerrazzanoModelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VerrazzanoModelList{ListMeta: obj.(*v1beta1.VerrazzanoModelList).ListMeta}
	for _, item := range obj.(*v1beta1.VerrazzanoModelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested verrazzanoModels.
func (c *FakeVerrazzanoModels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(verrazzanomodelsResource, c.ns, opts))

}

// Create takes the representation of a verrazzanoModel and creates it.  Returns the server's representation of the verrazzanoModel, and an error, if there is any.
func (c *FakeVerrazzanoModels) Create(verrazzanoModel *v1beta1.VerrazzanoModel) (result *v1beta1.VerrazzanoModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(verrazzanomodelsResource, c.ns, verrazzanoModel), &v1beta1.VerrazzanoModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VerrazzanoModel), err
}

// Update takes the representation of a verrazzanoModel and updates it. Returns the server's representation of the verrazzanoModel, and an error, if there is any.
func (c *FakeVerrazzanoModels) Update(verrazzanoModel *v1beta1.VerrazzanoModel) (result *v1beta1.VerrazzanoModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(verrazzanomodelsResource, c.ns, verrazzanoModel), &v1beta1.VerrazzanoModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VerrazzanoModel), err
}

// Delete takes name of the verrazzanoModel and deletes it. Returns an error if one occurs.
func (c *FakeVerrazzanoModels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(verrazzanomodelsResource, c.ns, name), &v1beta1.VerrazzanoModel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVerrazzanoModels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(verrazzanomodelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.VerrazzanoModelList{})
	return err
}

// Patch applies the patch and returns the patched verrazzanoModel.
func (c *FakeVerrazzanoModels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VerrazzanoModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(verrazzanomodelsResource, c.ns, name, pt, data, subresources...), &v1beta1.VerrazzanoModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VerrazzanoModel), err
}
