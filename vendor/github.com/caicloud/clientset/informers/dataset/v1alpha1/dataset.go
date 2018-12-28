/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/dataset/v1alpha1"
	datasetv1alpha1 "github.com/caicloud/clientset/pkg/apis/dataset/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// DatasetInformer provides access to a shared informer and lister for
// Datasets.
type DatasetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DatasetLister
}

type datasetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDatasetInformer constructs a new informer for Dataset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatasetInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatasetInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDatasetInformer constructs a new informer for Dataset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatasetInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DatasetV1alpha1().Datasets().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DatasetV1alpha1().Datasets().Watch(options)
			},
		},
		&datasetv1alpha1.Dataset{},
		resyncPeriod,
		indexers,
	)
}

func (f *datasetInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatasetInformer(client.(kubernetes.Interface), resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *datasetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&datasetv1alpha1.Dataset{}, f.defaultInformer)
}

func (f *datasetInformer) Lister() v1alpha1.DatasetLister {
	return v1alpha1.NewDatasetLister(f.Informer().GetIndexer())
}