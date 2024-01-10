// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v0alpha1 "github.com/grafana/grafana/pkg/apis/service/v0alpha1"
	servicev0alpha1 "github.com/grafana/grafana/pkg/generated/applyconfiguration/service/v0alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExternalNames implements ExternalNameInterface
type FakeExternalNames struct {
	Fake *FakeServiceV0alpha1
	ns   string
}

var externalnamesResource = v0alpha1.SchemeGroupVersion.WithResource("externalnames")

var externalnamesKind = v0alpha1.SchemeGroupVersion.WithKind("ExternalName")

// Get takes name of the externalName, and returns the corresponding externalName object, and an error if there is any.
func (c *FakeExternalNames) Get(ctx context.Context, name string, options v1.GetOptions) (result *v0alpha1.ExternalName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalnamesResource, c.ns, name), &v0alpha1.ExternalName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v0alpha1.ExternalName), err
}

// List takes label and field selectors, and returns the list of ExternalNames that match those selectors.
func (c *FakeExternalNames) List(ctx context.Context, opts v1.ListOptions) (result *v0alpha1.ExternalNameList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalnamesResource, externalnamesKind, c.ns, opts), &v0alpha1.ExternalNameList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v0alpha1.ExternalNameList{ListMeta: obj.(*v0alpha1.ExternalNameList).ListMeta}
	for _, item := range obj.(*v0alpha1.ExternalNameList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalNames.
func (c *FakeExternalNames) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalnamesResource, c.ns, opts))

}

// Create takes the representation of a externalName and creates it.  Returns the server's representation of the externalName, and an error, if there is any.
func (c *FakeExternalNames) Create(ctx context.Context, externalName *v0alpha1.ExternalName, opts v1.CreateOptions) (result *v0alpha1.ExternalName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalnamesResource, c.ns, externalName), &v0alpha1.ExternalName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v0alpha1.ExternalName), err
}

// Update takes the representation of a externalName and updates it. Returns the server's representation of the externalName, and an error, if there is any.
func (c *FakeExternalNames) Update(ctx context.Context, externalName *v0alpha1.ExternalName, opts v1.UpdateOptions) (result *v0alpha1.ExternalName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalnamesResource, c.ns, externalName), &v0alpha1.ExternalName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v0alpha1.ExternalName), err
}

// Delete takes name of the externalName and deletes it. Returns an error if one occurs.
func (c *FakeExternalNames) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(externalnamesResource, c.ns, name, opts), &v0alpha1.ExternalName{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalNames) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalnamesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v0alpha1.ExternalNameList{})
	return err
}

// Patch applies the patch and returns the patched externalName.
func (c *FakeExternalNames) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v0alpha1.ExternalName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalnamesResource, c.ns, name, pt, data, subresources...), &v0alpha1.ExternalName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v0alpha1.ExternalName), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied externalName.
func (c *FakeExternalNames) Apply(ctx context.Context, externalName *servicev0alpha1.ExternalNameApplyConfiguration, opts v1.ApplyOptions) (result *v0alpha1.ExternalName, err error) {
	if externalName == nil {
		return nil, fmt.Errorf("externalName provided to Apply must not be nil")
	}
	data, err := json.Marshal(externalName)
	if err != nil {
		return nil, err
	}
	name := externalName.Name
	if name == nil {
		return nil, fmt.Errorf("externalName.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalnamesResource, c.ns, *name, types.ApplyPatchType, data), &v0alpha1.ExternalName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v0alpha1.ExternalName), err
}
