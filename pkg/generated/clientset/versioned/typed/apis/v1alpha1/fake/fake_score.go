/*
Copyright 2022 The Arbiter Authors.

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

	v1alpha1 "github.com/kube-arbiter/arbiter/pkg/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScores implements ScoreInterface
type FakeScores struct {
	Fake *FakeArbiterV1alpha1
	ns   string
}

var scoresResource = schema.GroupVersionResource{Group: "arbiter.k8s.com.cn", Version: "v1alpha1", Resource: "scores"}

var scoresKind = schema.GroupVersionKind{Group: "arbiter.k8s.com.cn", Version: "v1alpha1", Kind: "Score"}

// Get takes name of the score, and returns the corresponding score object, and an error if there is any.
func (c *FakeScores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Score, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scoresResource, c.ns, name), &v1alpha1.Score{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Score), err
}

// List takes label and field selectors, and returns the list of Scores that match those selectors.
func (c *FakeScores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scoresResource, scoresKind, c.ns, opts), &v1alpha1.ScoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScoreList{ListMeta: obj.(*v1alpha1.ScoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scores.
func (c *FakeScores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scoresResource, c.ns, opts))

}

// Create takes the representation of a score and creates it.  Returns the server's representation of the score, and an error, if there is any.
func (c *FakeScores) Create(ctx context.Context, score *v1alpha1.Score, opts v1.CreateOptions) (result *v1alpha1.Score, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scoresResource, c.ns, score), &v1alpha1.Score{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Score), err
}

// Update takes the representation of a score and updates it. Returns the server's representation of the score, and an error, if there is any.
func (c *FakeScores) Update(ctx context.Context, score *v1alpha1.Score, opts v1.UpdateOptions) (result *v1alpha1.Score, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scoresResource, c.ns, score), &v1alpha1.Score{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Score), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScores) UpdateStatus(ctx context.Context, score *v1alpha1.Score, opts v1.UpdateOptions) (*v1alpha1.Score, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scoresResource, "status", c.ns, score), &v1alpha1.Score{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Score), err
}

// Delete takes name of the score and deletes it. Returns an error if one occurs.
func (c *FakeScores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scoresResource, c.ns, name, opts), &v1alpha1.Score{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScoreList{})
	return err
}

// Patch applies the patch and returns the patched score.
func (c *FakeScores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Score, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.Score{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Score), err
}
