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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// APIServiceExportResourceLister helps list APIServiceExportResources.
// All objects returned here must be treated as read-only.
type APIServiceExportResourceLister interface {
	// List lists all APIServiceExportResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.APIServiceExportResource, err error)
	// APIServiceExportResources returns an object that can list and get APIServiceExportResources.
	APIServiceExportResources(namespace string) APIServiceExportResourceNamespaceLister
	APIServiceExportResourceListerExpansion
}

// aPIServiceExportResourceLister implements the APIServiceExportResourceLister interface.
type aPIServiceExportResourceLister struct {
	indexer cache.Indexer
}

// NewAPIServiceExportResourceLister returns a new APIServiceExportResourceLister.
func NewAPIServiceExportResourceLister(indexer cache.Indexer) APIServiceExportResourceLister {
	return &aPIServiceExportResourceLister{indexer: indexer}
}

// List lists all APIServiceExportResources in the indexer.
func (s *aPIServiceExportResourceLister) List(selector labels.Selector) (ret []*v1alpha1.APIServiceExportResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIServiceExportResource))
	})
	return ret, err
}

// APIServiceExportResources returns an object that can list and get APIServiceExportResources.
func (s *aPIServiceExportResourceLister) APIServiceExportResources(namespace string) APIServiceExportResourceNamespaceLister {
	return aPIServiceExportResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// APIServiceExportResourceNamespaceLister helps list and get APIServiceExportResources.
// All objects returned here must be treated as read-only.
type APIServiceExportResourceNamespaceLister interface {
	// List lists all APIServiceExportResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.APIServiceExportResource, err error)
	// Get retrieves the APIServiceExportResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.APIServiceExportResource, error)
	APIServiceExportResourceNamespaceListerExpansion
}

// aPIServiceExportResourceNamespaceLister implements the APIServiceExportResourceNamespaceLister
// interface.
type aPIServiceExportResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all APIServiceExportResources in the indexer for a given namespace.
func (s aPIServiceExportResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.APIServiceExportResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIServiceExportResource))
	})
	return ret, err
}

// Get retrieves the APIServiceExportResource from the indexer for a given namespace and name.
func (s aPIServiceExportResourceNamespaceLister) Get(name string) (*v1alpha1.APIServiceExportResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apiserviceexportresource"), name)
	}
	return obj.(*v1alpha1.APIServiceExportResource), nil
}