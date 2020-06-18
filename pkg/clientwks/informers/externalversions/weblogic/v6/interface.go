// Copyright (c) 2020, Oracle Corporation and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by informer-gen. DO NOT EDIT.

package v6

import (
	internalinterfaces "github.com/verrazzano/verrazzano-crd-generator/pkg/clientwks/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Domains returns a DomainInformer.
	Domains() DomainInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Domains returns a DomainInformer.
func (v *version) Domains() DomainInformer {
	return &domainInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
