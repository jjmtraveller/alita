// Copyright 2019 The Alita Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	slurmalitaiov1alpha1 "github.com/alita/alita/pkg/apis/slurm.alita.io/v1alpha1"
	versioned "github.com/alita/alita/pkg/client/clientset/versioned"
	internalinterfaces "github.com/alita/alita/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/alita/alita/pkg/client/listers/slurm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterInformer provides access to a shared informer and lister for
// Clusters.
type ClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterLister
}

type clusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewClusterInformer constructs a new informer for Cluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.SlurmV1alpha1().Clusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.SlurmV1alpha1().Clusters(namespace).Watch(options)
			},
		},
		&slurmalitaiov1alpha1.Cluster{},
		resyncPeriod,
		indexers,
	)
}

func defaultClusterInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewClusterInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *clusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&slurmalitaiov1alpha1.Cluster{}, defaultClusterInformer)
}

func (f *clusterInformer) Lister() v1alpha1.ClusterLister {
	return v1alpha1.NewClusterLister(f.Informer().GetIndexer())
}