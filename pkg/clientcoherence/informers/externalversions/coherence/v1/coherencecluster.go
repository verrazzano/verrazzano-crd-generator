// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	coherencev1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/coherence/v1"
	versioned "github.com/verrazzano/verrazzano-crd-generator/pkg/clientcoherence/clientset/versioned"
	internalinterfaces "github.com/verrazzano/verrazzano-crd-generator/pkg/clientcoherence/informers/externalversions/internalinterfaces"
	v1 "github.com/verrazzano/verrazzano-crd-generator/pkg/clientcoherence/listers/coherence/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CoherenceClusterInformer provides access to a shared informer and lister for
// CoherenceClusters.
type CoherenceClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CoherenceClusterLister
}

type coherenceClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCoherenceClusterInformer constructs a new informer for CoherenceCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCoherenceClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCoherenceClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCoherenceClusterInformer constructs a new informer for CoherenceCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCoherenceClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoherenceV1().CoherenceClusters(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoherenceV1().CoherenceClusters(namespace).Watch(options)
			},
		},
		&coherencev1.CoherenceCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *coherenceClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCoherenceClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *coherenceClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&coherencev1.CoherenceCluster{}, f.defaultInformer)
}

func (f *coherenceClusterInformer) Lister() v1.CoherenceClusterLister {
	return v1.NewCoherenceClusterLister(f.Informer().GetIndexer())
}