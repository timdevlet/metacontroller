/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	metacontroller_v1alpha1 "k8s.io/metacontroller/apis/metacontroller/v1alpha1"
	internalclientset "k8s.io/metacontroller/client/generated/clientset/internalclientset"
	internalinterfaces "k8s.io/metacontroller/client/generated/informer/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/metacontroller/client/generated/lister/metacontroller/v1alpha1"
)

// ControllerRevisionInformer provides access to a shared informer and lister for
// ControllerRevisions.
type ControllerRevisionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ControllerRevisionLister
}

type controllerRevisionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewControllerRevisionInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredControllerRevisionInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetacontrollerV1alpha1().ControllerRevisions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetacontrollerV1alpha1().ControllerRevisions(namespace).Watch(options)
			},
		},
		&metacontroller_v1alpha1.ControllerRevision{},
		resyncPeriod,
		indexers,
	)
}

func (f *controllerRevisionInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&metacontroller_v1alpha1.ControllerRevision{}, f.defaultInformer)
}

func (f *controllerRevisionInformer) Lister() v1alpha1.ControllerRevisionLister {
	return v1alpha1.NewControllerRevisionLister(f.Informer().GetIndexer())
}
