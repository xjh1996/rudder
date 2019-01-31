/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/rbac/v1alpha1"
)

var _ v1alpha1.ClusterRoleLister = &clusterRoleLister{}

// clusterRoleLister implements the ClusterRoleLister interface.
type clusterRoleLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterRoleLister returns a new ClusterRoleLister.
func NewClusterRoleLister(client kubernetes.Interface) v1alpha1.ClusterRoleLister {
	return NewFilteredClusterRoleLister(client, nil)
}

func NewFilteredClusterRoleLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1alpha1.ClusterRoleLister {
	return &clusterRoleLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all ClusterRoles in the indexer.
func (s *clusterRoleLister) List(selector labels.Selector) (ret []*rbacv1alpha1.ClusterRole, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.RbacV1alpha1().ClusterRoles().List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the ClusterRole from the index for a given name.
func (s *clusterRoleLister) Get(name string) (*rbacv1alpha1.ClusterRole, error) {
	return s.client.RbacV1alpha1().ClusterRoles().Get(name, v1.GetOptions{})
}
