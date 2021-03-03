/*
Copyright 2021 The OpenEBS Authors

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
	"context"

	v1alpha1 "github.com/openebs/api/v2/pkg/apis/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBlockDeviceClaims implements BlockDeviceClaimInterface
type FakeBlockDeviceClaims struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var blockdeviceclaimsResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "blockdeviceclaims"}

var blockdeviceclaimsKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "BlockDeviceClaim"}

// Get takes name of the blockDeviceClaim, and returns the corresponding blockDeviceClaim object, and an error if there is any.
func (c *FakeBlockDeviceClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BlockDeviceClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(blockdeviceclaimsResource, c.ns, name), &v1alpha1.BlockDeviceClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlockDeviceClaim), err
}

// List takes label and field selectors, and returns the list of BlockDeviceClaims that match those selectors.
func (c *FakeBlockDeviceClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BlockDeviceClaimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(blockdeviceclaimsResource, blockdeviceclaimsKind, c.ns, opts), &v1alpha1.BlockDeviceClaimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BlockDeviceClaimList{ListMeta: obj.(*v1alpha1.BlockDeviceClaimList).ListMeta}
	for _, item := range obj.(*v1alpha1.BlockDeviceClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested blockDeviceClaims.
func (c *FakeBlockDeviceClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(blockdeviceclaimsResource, c.ns, opts))

}

// Create takes the representation of a blockDeviceClaim and creates it.  Returns the server's representation of the blockDeviceClaim, and an error, if there is any.
func (c *FakeBlockDeviceClaims) Create(ctx context.Context, blockDeviceClaim *v1alpha1.BlockDeviceClaim, opts v1.CreateOptions) (result *v1alpha1.BlockDeviceClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(blockdeviceclaimsResource, c.ns, blockDeviceClaim), &v1alpha1.BlockDeviceClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlockDeviceClaim), err
}

// Update takes the representation of a blockDeviceClaim and updates it. Returns the server's representation of the blockDeviceClaim, and an error, if there is any.
func (c *FakeBlockDeviceClaims) Update(ctx context.Context, blockDeviceClaim *v1alpha1.BlockDeviceClaim, opts v1.UpdateOptions) (result *v1alpha1.BlockDeviceClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(blockdeviceclaimsResource, c.ns, blockDeviceClaim), &v1alpha1.BlockDeviceClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlockDeviceClaim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBlockDeviceClaims) UpdateStatus(ctx context.Context, blockDeviceClaim *v1alpha1.BlockDeviceClaim, opts v1.UpdateOptions) (*v1alpha1.BlockDeviceClaim, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(blockdeviceclaimsResource, "status", c.ns, blockDeviceClaim), &v1alpha1.BlockDeviceClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlockDeviceClaim), err
}

// Delete takes name of the blockDeviceClaim and deletes it. Returns an error if one occurs.
func (c *FakeBlockDeviceClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(blockdeviceclaimsResource, c.ns, name), &v1alpha1.BlockDeviceClaim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBlockDeviceClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(blockdeviceclaimsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BlockDeviceClaimList{})
	return err
}

// Patch applies the patch and returns the patched blockDeviceClaim.
func (c *FakeBlockDeviceClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BlockDeviceClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(blockdeviceclaimsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BlockDeviceClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlockDeviceClaim), err
}
