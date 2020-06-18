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

// CoherenceClusterLister helps list CoherenceClusters.
type CoherenceClusterLister interface {
	// List lists all CoherenceClusters in the indexer.
	List(selector labels.Selector) (ret []*v1.CoherenceCluster, err error)
	// CoherenceClusters returns an object that can list and get CoherenceClusters.
	CoherenceClusters(namespace string) CoherenceClusterNamespaceLister
	CoherenceClusterListerExpansion
}

// coherenceClusterLister implements the CoherenceClusterLister interface.
type coherenceClusterLister struct {
	indexer cache.Indexer
}

// NewCoherenceClusterLister returns a new CoherenceClusterLister.
func NewCoherenceClusterLister(indexer cache.Indexer) CoherenceClusterLister {
	return &coherenceClusterLister{indexer: indexer}
}

// List lists all CoherenceClusters in the indexer.
func (s *coherenceClusterLister) List(selector labels.Selector) (ret []*v1.CoherenceCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CoherenceCluster))
	})
	return ret, err
}

// CoherenceClusters returns an object that can list and get CoherenceClusters.
func (s *coherenceClusterLister) CoherenceClusters(namespace string) CoherenceClusterNamespaceLister {
	return coherenceClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CoherenceClusterNamespaceLister helps list and get CoherenceClusters.
type CoherenceClusterNamespaceLister interface {
	// List lists all CoherenceClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CoherenceCluster, err error)
	// Get retrieves the CoherenceCluster from the indexer for a given namespace and name.
	Get(name string) (*v1.CoherenceCluster, error)
	CoherenceClusterNamespaceListerExpansion
}

// coherenceClusterNamespaceLister implements the CoherenceClusterNamespaceLister
// interface.
type coherenceClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CoherenceClusters in the indexer for a given namespace.
func (s coherenceClusterNamespaceLister) List(selector labels.Selector) (ret []*v1.CoherenceCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CoherenceCluster))
	})
	return ret, err
}

// Get retrieves the CoherenceCluster from the indexer for a given namespace and name.
func (s coherenceClusterNamespaceLister) Get(name string) (*v1.CoherenceCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("coherencecluster"), name)
	}
	return obj.(*v1.CoherenceCluster), nil
}