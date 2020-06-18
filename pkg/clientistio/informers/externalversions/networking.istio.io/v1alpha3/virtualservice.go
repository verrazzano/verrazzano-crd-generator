// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	time "time"

	networkingistioiov1alpha3 "github.com/verrazzano/verrazzano-crd-generator/pkg/apis/networking.istio.io/v1alpha3"
	versioned "github.com/verrazzano/verrazzano-crd-generator/pkg/clientistio/clientset/versioned"
	internalinterfaces "github.com/verrazzano/verrazzano-crd-generator/pkg/clientistio/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/verrazzano/verrazzano-crd-generator/pkg/clientistio/listers/networking.istio.io/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualServiceInformer provides access to a shared informer and lister for
// VirtualServices.
type VirtualServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.VirtualServiceLister
}

type virtualServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualServiceInformer constructs a new informer for VirtualService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualServiceInformer constructs a new informer for VirtualService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().VirtualServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().VirtualServices(namespace).Watch(options)
			},
		},
		&networkingistioiov1alpha3.VirtualService{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingistioiov1alpha3.VirtualService{}, f.defaultInformer)
}

func (f *virtualServiceInformer) Lister() v1alpha3.VirtualServiceLister {
	return v1alpha3.NewVirtualServiceLister(f.Informer().GetIndexer())
}
