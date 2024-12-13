//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kube Bind Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kubebindv1alpha1 "github.com/kube-bind/kube-bind/sdk/apis/kubebind/v1alpha1"
	scopedclientset "github.com/kube-bind/kube-bind/sdk/kcp/clientset/versioned"
	clientset "github.com/kube-bind/kube-bind/sdk/kcp/clientset/versioned/cluster"
	"github.com/kube-bind/kube-bind/sdk/kcp/informers/externalversions/internalinterfaces"
	kubebindv1alpha1listers "github.com/kube-bind/kube-bind/sdk/kcp/listers/kubebind/v1alpha1"
)

// APIServiceExportClusterInformer provides access to a shared informer and lister for
// APIServiceExports.
type APIServiceExportClusterInformer interface {
	Cluster(logicalcluster.Name) APIServiceExportInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() kubebindv1alpha1listers.APIServiceExportClusterLister
}

type aPIServiceExportClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIServiceExportClusterInformer constructs a new informer for APIServiceExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIServiceExportClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIServiceExportClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIServiceExportClusterInformer constructs a new informer for APIServiceExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIServiceExportClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().APIServiceExports().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().APIServiceExports().Watch(context.TODO(), options)
			},
		},
		&kubebindv1alpha1.APIServiceExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIServiceExportClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIServiceExportClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *aPIServiceExportClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&kubebindv1alpha1.APIServiceExport{}, f.defaultInformer)
}

func (f *aPIServiceExportClusterInformer) Lister() kubebindv1alpha1listers.APIServiceExportClusterLister {
	return kubebindv1alpha1listers.NewAPIServiceExportClusterLister(f.Informer().GetIndexer())
}

// APIServiceExportInformer provides access to a shared informer and lister for
// APIServiceExports.
type APIServiceExportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubebindv1alpha1listers.APIServiceExportLister
}

func (f *aPIServiceExportClusterInformer) Cluster(clusterName logicalcluster.Name) APIServiceExportInformer {
	return &aPIServiceExportInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type aPIServiceExportInformer struct {
	informer cache.SharedIndexInformer
	lister   kubebindv1alpha1listers.APIServiceExportLister
}

func (f *aPIServiceExportInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *aPIServiceExportInformer) Lister() kubebindv1alpha1listers.APIServiceExportLister {
	return f.lister
}

type aPIServiceExportScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func (f *aPIServiceExportScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubebindv1alpha1.APIServiceExport{}, f.defaultInformer)
}

func (f *aPIServiceExportScopedInformer) Lister() kubebindv1alpha1listers.APIServiceExportLister {
	return kubebindv1alpha1listers.NewAPIServiceExportLister(f.Informer().GetIndexer())
}

// NewAPIServiceExportInformer constructs a new informer for APIServiceExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIServiceExportInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIServiceExportInformer(client, resyncPeriod, namespace, indexers, nil)
}

// NewFilteredAPIServiceExportInformer constructs a new informer for APIServiceExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIServiceExportInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().APIServiceExports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().APIServiceExports(namespace).Watch(context.TODO(), options)
			},
		},
		&kubebindv1alpha1.APIServiceExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIServiceExportScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIServiceExportInformer(client, resyncPeriod, f.namespace, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	}, f.tweakListOptions)
}
