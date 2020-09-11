/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	cloudgatewayv1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
)

// FakeVirtualPresences implements VirtualPresenceInterface
type FakeVirtualPresences struct {
	Fake *FakeCloudgatewayV1
	ns   string
	te   string
}

var virtualpresencesResource = schema.GroupVersionResource{Group: "cloudgateway.arktos.io", Version: "v1", Resource: "virtualpresences"}

var virtualpresencesKind = schema.GroupVersionKind{Group: "cloudgateway.arktos.io", Version: "v1", Kind: "VirtualPresence"}

// Get takes name of the virtualPresence, and returns the corresponding virtualPresence object, and an error if there is any.
func (c *FakeVirtualPresences) Get(name string, options v1.GetOptions) (result *cloudgatewayv1.VirtualPresence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithMultiTenancy(virtualpresencesResource, c.ns, name, c.te), &cloudgatewayv1.VirtualPresence{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.VirtualPresence), err
}

// List takes label and field selectors, and returns the list of VirtualPresences that match those selectors.
func (c *FakeVirtualPresences) List(opts v1.ListOptions) (result *cloudgatewayv1.VirtualPresenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithMultiTenancy(virtualpresencesResource, virtualpresencesKind, c.ns, opts, c.te), &cloudgatewayv1.VirtualPresenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudgatewayv1.VirtualPresenceList{ListMeta: obj.(*cloudgatewayv1.VirtualPresenceList).ListMeta}
	for _, item := range obj.(*cloudgatewayv1.VirtualPresenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested virtualPresences.
func (c *FakeVirtualPresences) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewWatchActionWithMultiTenancy(virtualpresencesResource, c.ns, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a virtualPresence and creates it.  Returns the server's representation of the virtualPresence, and an error, if there is any.
func (c *FakeVirtualPresences) Create(virtualPresence *cloudgatewayv1.VirtualPresence) (result *cloudgatewayv1.VirtualPresence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithMultiTenancy(virtualpresencesResource, c.ns, virtualPresence, c.te), &cloudgatewayv1.VirtualPresence{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.VirtualPresence), err
}

// Update takes the representation of a virtualPresence and updates it. Returns the server's representation of the virtualPresence, and an error, if there is any.
func (c *FakeVirtualPresences) Update(virtualPresence *cloudgatewayv1.VirtualPresence) (result *cloudgatewayv1.VirtualPresence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithMultiTenancy(virtualpresencesResource, c.ns, virtualPresence, c.te), &cloudgatewayv1.VirtualPresence{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.VirtualPresence), err
}

// Delete takes name of the virtualPresence and deletes it. Returns an error if one occurs.
func (c *FakeVirtualPresences) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithMultiTenancy(virtualpresencesResource, c.ns, name, c.te), &cloudgatewayv1.VirtualPresence{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualPresences) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithMultiTenancy(virtualpresencesResource, c.ns, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &cloudgatewayv1.VirtualPresenceList{})
	return err
}

// Patch applies the patch and returns the patched virtualPresence.
func (c *FakeVirtualPresences) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudgatewayv1.VirtualPresence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithMultiTenancy(virtualpresencesResource, c.te, c.ns, name, pt, data, subresources...), &cloudgatewayv1.VirtualPresence{})

	if obj == nil {
		return nil, err
	}

	return obj.(*cloudgatewayv1.VirtualPresence), err
}