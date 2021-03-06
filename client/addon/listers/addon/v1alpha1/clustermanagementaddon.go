// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/open-cluster-management/api/addon/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterManagementAddOnLister helps list ClusterManagementAddOns.
// All objects returned here must be treated as read-only.
type ClusterManagementAddOnLister interface {
	// List lists all ClusterManagementAddOns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterManagementAddOn, err error)
	// Get retrieves the ClusterManagementAddOn from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterManagementAddOn, error)
	ClusterManagementAddOnListerExpansion
}

// clusterManagementAddOnLister implements the ClusterManagementAddOnLister interface.
type clusterManagementAddOnLister struct {
	indexer cache.Indexer
}

// NewClusterManagementAddOnLister returns a new ClusterManagementAddOnLister.
func NewClusterManagementAddOnLister(indexer cache.Indexer) ClusterManagementAddOnLister {
	return &clusterManagementAddOnLister{indexer: indexer}
}

// List lists all ClusterManagementAddOns in the indexer.
func (s *clusterManagementAddOnLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterManagementAddOn, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterManagementAddOn))
	})
	return ret, err
}

// Get retrieves the ClusterManagementAddOn from the index for a given name.
func (s *clusterManagementAddOnLister) Get(name string) (*v1alpha1.ClusterManagementAddOn, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustermanagementaddon"), name)
	}
	return obj.(*v1alpha1.ClusterManagementAddOn), nil
}
