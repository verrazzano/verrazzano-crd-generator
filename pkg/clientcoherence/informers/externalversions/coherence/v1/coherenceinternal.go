// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	coherencev1 "github.com/oracle/verrazzano-crd-generator/pkg/apis/coherence/v1"
	versioned "github.com/oracle/verrazzano-crd-generator/pkg/clientcoherence/clientset/versioned"
	internalinterfaces "github.com/oracle/verrazzano-crd-generator/pkg/clientcoherence/informers/externalversions/internalinterfaces"
	v1 "github.com/oracle/verrazzano-crd-generator/pkg/clientcoherence/listers/coherence/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CoherenceInternalInformer provides access to a shared informer and lister for
// CoherenceInternals.
type CoherenceInternalInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CoherenceInternalLister
}

type coherenceInternalInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCoherenceInternalInformer constructs a new informer for CoherenceInternal type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCoherenceInternalInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCoherenceInternalInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCoherenceInternalInformer constructs a new informer for CoherenceInternal type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCoherenceInternalInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoherenceV1().CoherenceInternals(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoherenceV1().CoherenceInternals(namespace).Watch(options)
			},
		},
		&coherencev1.CoherenceInternal{},
		resyncPeriod,
		indexers,
	)
}

func (f *coherenceInternalInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCoherenceInternalInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *coherenceInternalInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&coherencev1.CoherenceInternal{}, f.defaultInformer)
}

func (f *coherenceInternalInformer) Lister() v1.CoherenceInternalLister {
	return v1.NewCoherenceInternalLister(f.Informer().GetIndexer())
}
