/*
Copyright 2018 Giant Swarm GmbH.

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
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CertConfigsGetter has a method to return a CertConfigInterface.
// A group's client should implement this interface.
type CertConfigsGetter interface {
	CertConfigs(namespace string) CertConfigInterface
}

// CertConfigInterface has methods to work with CertConfig resources.
type CertConfigInterface interface {
	Create(*v1alpha1.CertConfig) (*v1alpha1.CertConfig, error)
	Update(*v1alpha1.CertConfig) (*v1alpha1.CertConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CertConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.CertConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CertConfig, err error)
	CertConfigExpansion
}

// certConfigs implements CertConfigInterface
type certConfigs struct {
	client rest.Interface
	ns     string
}

// newCertConfigs returns a CertConfigs
func newCertConfigs(c *CoreV1alpha1Client, namespace string) *certConfigs {
	return &certConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the certConfig, and returns the corresponding certConfig object, and an error if there is any.
func (c *certConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CertConfig, err error) {
	result = &v1alpha1.CertConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CertConfigs that match those selectors.
func (c *certConfigs) List(opts v1.ListOptions) (result *v1alpha1.CertConfigList, err error) {
	result = &v1alpha1.CertConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certConfigs.
func (c *certConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("certconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a certConfig and creates it.  Returns the server's representation of the certConfig, and an error, if there is any.
func (c *certConfigs) Create(certConfig *v1alpha1.CertConfig) (result *v1alpha1.CertConfig, err error) {
	result = &v1alpha1.CertConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("certconfigs").
		Body(certConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a certConfig and updates it. Returns the server's representation of the certConfig, and an error, if there is any.
func (c *certConfigs) Update(certConfig *v1alpha1.CertConfig) (result *v1alpha1.CertConfig, err error) {
	result = &v1alpha1.CertConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certconfigs").
		Name(certConfig.Name).
		Body(certConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the certConfig and deletes it. Returns an error if one occurs.
func (c *certConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched certConfig.
func (c *certConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CertConfig, err error) {
	result = &v1alpha1.CertConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("certconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
