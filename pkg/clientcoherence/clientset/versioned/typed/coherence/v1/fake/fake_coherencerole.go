// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	coherencev1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/coherence/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCoherenceRoles implements CoherenceRoleInterface
type FakeCoherenceRoles struct {
	Fake *FakeCoherenceV1
	ns   string
}

var coherencerolesResource = schema.GroupVersionResource{Group: "coherence.oracle.com", Version: "v1", Resource: "coherenceroles"}

var coherencerolesKind = schema.GroupVersionKind{Group: "coherence.oracle.com", Version: "v1", Kind: "CoherenceRole"}

// Get takes name of the coherenceRole, and returns the corresponding coherenceRole object, and an error if there is any.
func (c *FakeCoherenceRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *coherencev1.CoherenceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(coherencerolesResource, c.ns, name), &coherencev1.CoherenceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coherencev1.CoherenceRole), err
}

// List takes label and field selectors, and returns the list of CoherenceRoles that match those selectors.
func (c *FakeCoherenceRoles) List(ctx context.Context, opts v1.ListOptions) (result *coherencev1.CoherenceRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(coherencerolesResource, coherencerolesKind, c.ns, opts), &coherencev1.CoherenceRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &coherencev1.CoherenceRoleList{ListMeta: obj.(*coherencev1.CoherenceRoleList).ListMeta}
	for _, item := range obj.(*coherencev1.CoherenceRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested coherenceRoles.
func (c *FakeCoherenceRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(coherencerolesResource, c.ns, opts))

}

// Create takes the representation of a coherenceRole and creates it.  Returns the server's representation of the coherenceRole, and an error, if there is any.
func (c *FakeCoherenceRoles) Create(ctx context.Context, coherenceRole *coherencev1.CoherenceRole, opts v1.CreateOptions) (result *coherencev1.CoherenceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(coherencerolesResource, c.ns, coherenceRole), &coherencev1.CoherenceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coherencev1.CoherenceRole), err
}

// Update takes the representation of a coherenceRole and updates it. Returns the server's representation of the coherenceRole, and an error, if there is any.
func (c *FakeCoherenceRoles) Update(ctx context.Context, coherenceRole *coherencev1.CoherenceRole, opts v1.UpdateOptions) (result *coherencev1.CoherenceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(coherencerolesResource, c.ns, coherenceRole), &coherencev1.CoherenceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coherencev1.CoherenceRole), err
}

// Delete takes name of the coherenceRole and deletes it. Returns an error if one occurs.
func (c *FakeCoherenceRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(coherencerolesResource, c.ns, name), &coherencev1.CoherenceRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCoherenceRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(coherencerolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &coherencev1.CoherenceRoleList{})
	return err
}

// Patch applies the patch and returns the patched coherenceRole.
func (c *FakeCoherenceRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *coherencev1.CoherenceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(coherencerolesResource, c.ns, name, pt, data, subresources...), &coherencev1.CoherenceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coherencev1.CoherenceRole), err
}
