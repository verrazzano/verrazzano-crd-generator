// Copyright (c) 2020, Oracle and/or its affiliates.
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

// VerrazzanoBindingInformer provides access to a shared informer and lister for
// VerrazzanoBindings.
type VerrazzanoBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.VerrazzanoBindingLister
}

type verrazzanoBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVerrazzanoBindingInformer constructs a new informer for VerrazzanoBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVerrazzanoBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVerrazzanoBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVerrazzanoBindingInformer constructs a new informer for VerrazzanoBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVerrazzanoBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VerrazzanoV1beta1().VerrazzanoBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VerrazzanoV1beta1().VerrazzanoBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&verrazzanov1beta1.VerrazzanoBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *verrazzanoBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVerrazzanoBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *verrazzanoBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&verrazzanov1beta1.VerrazzanoBinding{}, f.defaultInformer)
}

func (f *verrazzanoBindingInformer) Lister() v1beta1.VerrazzanoBindingLister {
	return v1beta1.NewVerrazzanoBindingLister(f.Informer().GetIndexer())
}
