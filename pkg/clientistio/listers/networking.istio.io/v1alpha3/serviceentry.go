// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/oracle/verrazzano-crd-generator/pkg/apis/networking.istio.io/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceEntryLister helps list ServiceEntries.
type ServiceEntryLister interface {
	// List lists all ServiceEntries in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error)
	// ServiceEntries returns an object that can list and get ServiceEntries.
	ServiceEntries(namespace string) ServiceEntryNamespaceLister
	ServiceEntryListerExpansion
}

// serviceEntryLister implements the ServiceEntryLister interface.
type serviceEntryLister struct {
	indexer cache.Indexer
}

// NewServiceEntryLister returns a new ServiceEntryLister.
func NewServiceEntryLister(indexer cache.Indexer) ServiceEntryLister {
	return &serviceEntryLister{indexer: indexer}
}

// List lists all ServiceEntries in the indexer.
func (s *serviceEntryLister) List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ServiceEntry))
	})
	return ret, err
}

// ServiceEntries returns an object that can list and get ServiceEntries.
func (s *serviceEntryLister) ServiceEntries(namespace string) ServiceEntryNamespaceLister {
	return serviceEntryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceEntryNamespaceLister helps list and get ServiceEntries.
type ServiceEntryNamespaceLister interface {
	// List lists all ServiceEntries in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error)
	// Get retrieves the ServiceEntry from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.ServiceEntry, error)
	ServiceEntryNamespaceListerExpansion
}

// serviceEntryNamespaceLister implements the ServiceEntryNamespaceLister
// interface.
type serviceEntryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceEntries in the indexer for a given namespace.
func (s serviceEntryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ServiceEntry))
	})
	return ret, err
}

// Get retrieves the ServiceEntry from the indexer for a given namespace and name.
func (s serviceEntryNamespaceLister) Get(name string) (*v1alpha3.ServiceEntry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("serviceentry"), name)
	}
	return obj.(*v1alpha3.ServiceEntry), nil
}
