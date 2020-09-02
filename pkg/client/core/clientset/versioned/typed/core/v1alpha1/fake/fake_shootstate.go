/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShootStates implements ShootStateInterface
type FakeShootStates struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var shootstatesResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "v1alpha1", Resource: "shootstates"}

var shootstatesKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "v1alpha1", Kind: "ShootState"}

// Get takes name of the shootState, and returns the corresponding shootState object, and an error if there is any.
func (c *FakeShootStates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShootState, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shootstatesResource, c.ns, name), &v1alpha1.ShootState{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShootState), err
}

// List takes label and field selectors, and returns the list of ShootStates that match those selectors.
func (c *FakeShootStates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShootStateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shootstatesResource, shootstatesKind, c.ns, opts), &v1alpha1.ShootStateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShootStateList{ListMeta: obj.(*v1alpha1.ShootStateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShootStateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shootStates.
func (c *FakeShootStates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shootstatesResource, c.ns, opts))

}

// Create takes the representation of a shootState and creates it.  Returns the server's representation of the shootState, and an error, if there is any.
func (c *FakeShootStates) Create(ctx context.Context, shootState *v1alpha1.ShootState, opts v1.CreateOptions) (result *v1alpha1.ShootState, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shootstatesResource, c.ns, shootState), &v1alpha1.ShootState{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShootState), err
}

// Update takes the representation of a shootState and updates it. Returns the server's representation of the shootState, and an error, if there is any.
func (c *FakeShootStates) Update(ctx context.Context, shootState *v1alpha1.ShootState, opts v1.UpdateOptions) (result *v1alpha1.ShootState, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shootstatesResource, c.ns, shootState), &v1alpha1.ShootState{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShootState), err
}

// Delete takes name of the shootState and deletes it. Returns an error if one occurs.
func (c *FakeShootStates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shootstatesResource, c.ns, name), &v1alpha1.ShootState{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShootStates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shootstatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShootStateList{})
	return err
}

// Patch applies the patch and returns the patched shootState.
func (c *FakeShootStates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShootState, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shootstatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShootState{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShootState), err
}
