// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kong-operator/pkg/apis/kong/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KongLister helps list Kongs.
type KongLister interface {
	// List lists all Kongs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Kong, err error)
	// Kongs returns an object that can list and get Kongs.
	Kongs(namespace string) KongNamespaceLister
	KongListerExpansion
}

// kongLister implements the KongLister interface.
type kongLister struct {
	indexer cache.Indexer
}

// NewKongLister returns a new KongLister.
func NewKongLister(indexer cache.Indexer) KongLister {
	return &kongLister{indexer: indexer}
}

// List lists all Kongs in the indexer.
func (s *kongLister) List(selector labels.Selector) (ret []*v1alpha1.Kong, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Kong))
	})
	return ret, err
}

// Kongs returns an object that can list and get Kongs.
func (s *kongLister) Kongs(namespace string) KongNamespaceLister {
	return kongNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KongNamespaceLister helps list and get Kongs.
type KongNamespaceLister interface {
	// List lists all Kongs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Kong, err error)
	// Get retrieves the Kong from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Kong, error)
	KongNamespaceListerExpansion
}

// kongNamespaceLister implements the KongNamespaceLister
// interface.
type kongNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Kongs in the indexer for a given namespace.
func (s kongNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Kong, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Kong))
	})
	return ret, err
}

// Get retrieves the Kong from the indexer for a given namespace and name.
func (s kongNamespaceLister) Get(name string) (*v1alpha1.Kong, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kong"), name)
	}
	return obj.(*v1alpha1.Kong), nil
}
