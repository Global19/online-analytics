// This file was automatically generated by lister-gen

package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterRoleBindingLister helps list ClusterRoleBindings.
type ClusterRoleBindingLister interface {
	// List lists all ClusterRoleBindings in the indexer.
	List(selector labels.Selector) (ret []*authorization.ClusterRoleBinding, err error)
	// Get retrieves the ClusterRoleBinding from the index for a given name.
	Get(name string) (*authorization.ClusterRoleBinding, error)
	ClusterRoleBindingListerExpansion
}

// clusterRoleBindingLister implements the ClusterRoleBindingLister interface.
type clusterRoleBindingLister struct {
	indexer cache.Indexer
}

// NewClusterRoleBindingLister returns a new ClusterRoleBindingLister.
func NewClusterRoleBindingLister(indexer cache.Indexer) ClusterRoleBindingLister {
	return &clusterRoleBindingLister{indexer: indexer}
}

// List lists all ClusterRoleBindings in the indexer.
func (s *clusterRoleBindingLister) List(selector labels.Selector) (ret []*authorization.ClusterRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*authorization.ClusterRoleBinding))
	})
	return ret, err
}

// Get retrieves the ClusterRoleBinding from the index for a given name.
func (s *clusterRoleBindingLister) Get(name string) (*authorization.ClusterRoleBinding, error) {
	key := &authorization.ClusterRoleBinding{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authorization.Resource("clusterrolebinding"), name)
	}
	return obj.(*authorization.ClusterRoleBinding), nil
}