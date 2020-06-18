// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/coherence/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CoherenceRoleLister helps list CoherenceRoles.
type CoherenceRoleLister interface {
	// List lists all CoherenceRoles in the indexer.
	List(selector labels.Selector) (ret []*v1.CoherenceRole, err error)
	// CoherenceRoles returns an object that can list and get CoherenceRoles.
	CoherenceRoles(namespace string) CoherenceRoleNamespaceLister
	CoherenceRoleListerExpansion
}

// coherenceRoleLister implements the CoherenceRoleLister interface.
type coherenceRoleLister struct {
	indexer cache.Indexer
}

// NewCoherenceRoleLister returns a new CoherenceRoleLister.
func NewCoherenceRoleLister(indexer cache.Indexer) CoherenceRoleLister {
	return &coherenceRoleLister{indexer: indexer}
}

// List lists all CoherenceRoles in the indexer.
func (s *coherenceRoleLister) List(selector labels.Selector) (ret []*v1.CoherenceRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CoherenceRole))
	})
	return ret, err
}

// CoherenceRoles returns an object that can list and get CoherenceRoles.
func (s *coherenceRoleLister) CoherenceRoles(namespace string) CoherenceRoleNamespaceLister {
	return coherenceRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CoherenceRoleNamespaceLister helps list and get CoherenceRoles.
type CoherenceRoleNamespaceLister interface {
	// List lists all CoherenceRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CoherenceRole, err error)
	// Get retrieves the CoherenceRole from the indexer for a given namespace and name.
	Get(name string) (*v1.CoherenceRole, error)
	CoherenceRoleNamespaceListerExpansion
}

// coherenceRoleNamespaceLister implements the CoherenceRoleNamespaceLister
// interface.
type coherenceRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CoherenceRoles in the indexer for a given namespace.
func (s coherenceRoleNamespaceLister) List(selector labels.Selector) (ret []*v1.CoherenceRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CoherenceRole))
	})
	return ret, err
}

// Get retrieves the CoherenceRole from the indexer for a given namespace and name.
func (s coherenceRoleNamespaceLister) Get(name string) (*v1.CoherenceRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("coherencerole"), name)
	}
	return obj.(*v1.CoherenceRole), nil
}