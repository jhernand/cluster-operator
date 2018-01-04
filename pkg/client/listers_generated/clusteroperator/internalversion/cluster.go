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

// This file was automatically generated by lister-gen

package internalversion

import (
	clusteroperator "github.com/openshift/cluster-operator/pkg/apis/clusteroperator"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterLister helps list Clusters.
type ClusterLister interface {
	// List lists all Clusters in the indexer.
	List(selector labels.Selector) (ret []*clusteroperator.Cluster, err error)
	// Clusters returns an object that can list and get Clusters.
	Clusters(namespace string) ClusterNamespaceLister
	ClusterListerExpansion
}

// clusterLister implements the ClusterLister interface.
type clusterLister struct {
	indexer cache.Indexer
}

// NewClusterLister returns a new ClusterLister.
func NewClusterLister(indexer cache.Indexer) ClusterLister {
	return &clusterLister{indexer: indexer}
}

// List lists all Clusters in the indexer.
func (s *clusterLister) List(selector labels.Selector) (ret []*clusteroperator.Cluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*clusteroperator.Cluster))
	})
	return ret, err
}

// Clusters returns an object that can list and get Clusters.
func (s *clusterLister) Clusters(namespace string) ClusterNamespaceLister {
	return clusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterNamespaceLister helps list and get Clusters.
type ClusterNamespaceLister interface {
	// List lists all Clusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*clusteroperator.Cluster, err error)
	// Get retrieves the Cluster from the indexer for a given namespace and name.
	Get(name string) (*clusteroperator.Cluster, error)
	ClusterNamespaceListerExpansion
}

// clusterNamespaceLister implements the ClusterNamespaceLister
// interface.
type clusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Clusters in the indexer for a given namespace.
func (s clusterNamespaceLister) List(selector labels.Selector) (ret []*clusteroperator.Cluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*clusteroperator.Cluster))
	})
	return ret, err
}

// Get retrieves the Cluster from the indexer for a given namespace and name.
func (s clusterNamespaceLister) Get(name string) (*clusteroperator.Cluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(clusteroperator.Resource("cluster"), name)
	}
	return obj.(*clusteroperator.Cluster), nil
}
