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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/about-api/pkg/apis/v1alpha1"
	scheme "sigs.k8s.io/about-api/pkg/generated/clientset/versioned/scheme"
)

// ClusterPropertiesGetter has a method to return a ClusterPropertyInterface.
// A group's client should implement this interface.
type ClusterPropertiesGetter interface {
	ClusterProperties(namespace string) ClusterPropertyInterface
}

// ClusterPropertyInterface has methods to work with ClusterProperty resources.
type ClusterPropertyInterface interface {
	Create(ctx context.Context, clusterProperty *v1alpha1.ClusterProperty, opts v1.CreateOptions) (*v1alpha1.ClusterProperty, error)
	Update(ctx context.Context, clusterProperty *v1alpha1.ClusterProperty, opts v1.UpdateOptions) (*v1alpha1.ClusterProperty, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterProperty, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterPropertyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterProperty, err error)
	ClusterPropertyExpansion
}

// clusterProperties implements ClusterPropertyInterface
type clusterProperties struct {
	client rest.Interface
	ns     string
}

// newClusterProperties returns a ClusterProperties
func newClusterProperties(c *AboutV1alpha1Client, namespace string) *clusterProperties {
	return &clusterProperties{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterProperty, and returns the corresponding clusterProperty object, and an error if there is any.
func (c *clusterProperties) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterProperty, err error) {
	result = &v1alpha1.ClusterProperty{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterproperties").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterProperties that match those selectors.
func (c *clusterProperties) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPropertyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterPropertyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterproperties").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterProperties.
func (c *clusterProperties) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterproperties").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterProperty and creates it.  Returns the server's representation of the clusterProperty, and an error, if there is any.
func (c *clusterProperties) Create(ctx context.Context, clusterProperty *v1alpha1.ClusterProperty, opts v1.CreateOptions) (result *v1alpha1.ClusterProperty, err error) {
	result = &v1alpha1.ClusterProperty{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterproperties").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterProperty).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterProperty and updates it. Returns the server's representation of the clusterProperty, and an error, if there is any.
func (c *clusterProperties) Update(ctx context.Context, clusterProperty *v1alpha1.ClusterProperty, opts v1.UpdateOptions) (result *v1alpha1.ClusterProperty, err error) {
	result = &v1alpha1.ClusterProperty{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterproperties").
		Name(clusterProperty.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterProperty).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterProperty and deletes it. Returns an error if one occurs.
func (c *clusterProperties) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterproperties").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterProperties) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterproperties").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterProperty.
func (c *clusterProperties) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterProperty, err error) {
	result = &v1alpha1.ClusterProperty{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterproperties").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}