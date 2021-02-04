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

// VerrazzanoManagedClusterLister helps list VerrazzanoManagedClusters.
type VerrazzanoManagedClusterLister interface {
	// List lists all VerrazzanoManagedClusters in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VerrazzanoManagedCluster, err error)
	// VerrazzanoManagedClusters returns an object that can list and get VerrazzanoManagedClusters.
	VerrazzanoManagedClusters(namespace string) VerrazzanoManagedClusterNamespaceLister
	VerrazzanoManagedClusterListerExpansion
}

// verrazzanoManagedClusterLister implements the VerrazzanoManagedClusterLister interface.
type verrazzanoManagedClusterLister struct {
	indexer cache.Indexer
}

// NewVerrazzanoManagedClusterLister returns a new VerrazzanoManagedClusterLister.
func NewVerrazzanoManagedClusterLister(indexer cache.Indexer) VerrazzanoManagedClusterLister {
	return &verrazzanoManagedClusterLister{indexer: indexer}
}

// List lists all VerrazzanoManagedClusters in the indexer.
func (s *verrazzanoManagedClusterLister) List(selector labels.Selector) (ret []*v1beta1.VerrazzanoManagedCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VerrazzanoManagedCluster))
	})
	return ret, err
}

// VerrazzanoManagedClusters returns an object that can list and get VerrazzanoManagedClusters.
func (s *verrazzanoManagedClusterLister) VerrazzanoManagedClusters(namespace string) VerrazzanoManagedClusterNamespaceLister {
	return verrazzanoManagedClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VerrazzanoManagedClusterNamespaceLister helps list and get VerrazzanoManagedClusters.
type VerrazzanoManagedClusterNamespaceLister interface {
	// List lists all VerrazzanoManagedClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.VerrazzanoManagedCluster, err error)
	// Get retrieves the VerrazzanoManagedCluster from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.VerrazzanoManagedCluster, error)
	VerrazzanoManagedClusterNamespaceListerExpansion
}

// verrazzanoManagedClusterNamespaceLister implements the VerrazzanoManagedClusterNamespaceLister
// interface.
type verrazzanoManagedClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VerrazzanoManagedClusters in the indexer for a given namespace.
func (s verrazzanoManagedClusterNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VerrazzanoManagedCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VerrazzanoManagedCluster))
	})
	return ret, err
}

// Get retrieves the VerrazzanoManagedCluster from the indexer for a given namespace and name.
func (s verrazzanoManagedClusterNamespaceLister) Get(name string) (*v1beta1.VerrazzanoManagedCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("verrazzanomanagedcluster"), name)
	}
	return obj.(*v1beta1.VerrazzanoManagedCluster), nil
}
