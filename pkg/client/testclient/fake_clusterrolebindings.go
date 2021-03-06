package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/testing/core"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

// FakeClusterRoleBindings implements ClusterRoleBindingInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeClusterRoleBindings struct {
	Fake *Fake
}

var clusterRoleBindingsResource = unversioned.GroupVersionResource{Group: "", Version: "", Resource: "clusterrolebindings"}

func (c *FakeClusterRoleBindings) Get(name string) (*authorizationapi.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(core.NewRootGetAction(clusterRoleBindingsResource, name), &authorizationapi.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRoleBinding), err
}

func (c *FakeClusterRoleBindings) List(opts kapi.ListOptions) (*authorizationapi.ClusterRoleBindingList, error) {
	obj, err := c.Fake.Invokes(core.NewRootListAction(clusterRoleBindingsResource, opts), &authorizationapi.ClusterRoleBindingList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRoleBindingList), err
}

func (c *FakeClusterRoleBindings) Create(inObj *authorizationapi.ClusterRoleBinding) (*authorizationapi.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(clusterRoleBindingsResource, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRoleBinding), err
}

func (c *FakeClusterRoleBindings) Update(inObj *authorizationapi.ClusterRoleBinding) (*authorizationapi.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(core.NewRootUpdateAction(clusterRoleBindingsResource, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRoleBinding), err
}

func (c *FakeClusterRoleBindings) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(clusterRoleBindingsResource, name), &authorizationapi.ClusterRoleBinding{})
	return err
}
