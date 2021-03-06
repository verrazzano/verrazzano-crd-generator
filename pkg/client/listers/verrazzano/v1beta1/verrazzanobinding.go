// Copyright (c) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/verrazzano/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VerrazzanoBindingLister helps list VerrazzanoBindings.
type VerrazzanoBindingLister interface {
	// List lists all VerrazzanoBindings in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VerrazzanoBinding, err error)
	// VerrazzanoBindings returns an object that can list and get VerrazzanoBindings.
	VerrazzanoBindings(namespace string) VerrazzanoBindingNamespaceLister
	VerrazzanoBindingListerExpansion
}

// verrazzanoBindingLister implements the VerrazzanoBindingLister interface.
type verrazzanoBindingLister struct {
	indexer cache.Indexer
}

// NewVerrazzanoBindingLister returns a new VerrazzanoBindingLister.
func NewVerrazzanoBindingLister(indexer cache.Indexer) VerrazzanoBindingLister {
	return &verrazzanoBindingLister{indexer: indexer}
}

// List lists all VerrazzanoBindings in the indexer.
func (s *verrazzanoBindingLister) List(selector labels.Selector) (ret []*v1beta1.VerrazzanoBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VerrazzanoBinding))
	})
	return ret, err
}

// VerrazzanoBindings returns an object that can list and get VerrazzanoBindings.
func (s *verrazzanoBindingLister) VerrazzanoBindings(namespace string) VerrazzanoBindingNamespaceLister {
	return verrazzanoBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VerrazzanoBindingNamespaceLister helps list and get VerrazzanoBindings.
type VerrazzanoBindingNamespaceLister interface {
	// List lists all VerrazzanoBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.VerrazzanoBinding, err error)
	// Get retrieves the VerrazzanoBinding from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.VerrazzanoBinding, error)
	VerrazzanoBindingNamespaceListerExpansion
}

// verrazzanoBindingNamespaceLister implements the VerrazzanoBindingNamespaceLister
// interface.
type verrazzanoBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VerrazzanoBindings in the indexer for a given namespace.
func (s verrazzanoBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VerrazzanoBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VerrazzanoBinding))
	})
	return ret, err
}

// Get retrieves the VerrazzanoBinding from the indexer for a given namespace and name.
func (s verrazzanoBindingNamespaceLister) Get(name string) (*v1beta1.VerrazzanoBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("verrazzanobinding"), name)
	}
	return obj.(*v1beta1.VerrazzanoBinding), nil
}
