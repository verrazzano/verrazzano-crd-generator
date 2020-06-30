// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	verrazzanov1beta1 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/verrazzano/v1beta1"
	versioned "github.com/verrazzano/verrazzano-crd-generator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/verrazzano/verrazzano-crd-generator/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/verrazzano/verrazzano-crd-generator/pkg/client/listers/verrazzano/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VerrazzanoManagedClusterInformer provides access to a shared informer and lister for
// VerrazzanoManagedClusters.
type VerrazzanoManagedClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.VerrazzanoManagedClusterLister
}

type verrazzanoManagedClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVerrazzanoManagedClusterInformer constructs a new informer for VerrazzanoManagedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVerrazzanoManagedClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVerrazzanoManagedClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVerrazzanoManagedClusterInformer constructs a new informer for VerrazzanoManagedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVerrazzanoManagedClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VerrazzanoV1beta1().VerrazzanoManagedClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VerrazzanoV1beta1().VerrazzanoManagedClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&verrazzanov1beta1.VerrazzanoManagedCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *verrazzanoManagedClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVerrazzanoManagedClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *verrazzanoManagedClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&verrazzanov1beta1.VerrazzanoManagedCluster{}, f.defaultInformer)
}

func (f *verrazzanoManagedClusterInformer) Lister() v1beta1.VerrazzanoManagedClusterLister {
	return v1beta1.NewVerrazzanoManagedClusterLister(f.Informer().GetIndexer())
}
