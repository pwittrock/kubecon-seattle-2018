/*
Copyright 2018 The Kubernetes Authors.

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

package kubecon_seattle_2018


////////////////////////////////////////////////////////////////////
// Reconcile Request
////////////////////////////////////////////////////////////////////

// Sample Controller - Get Object Key
namespace, name, err := cache.SplitMetaNamespaceKey(key)
if err != nil {
	runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
	return nil
}

// Controller-Runtime - Get Object Key
namespace := request.Namespace
name := request.Name

////////////////////////////////////////////////////////////////////
// Client
////////////////////////////////////////////////////////////////////

// Sample Controller - Lookup Object
foo, err := c.foosLister.Foos(namespace).Get(name)

// Controller-Runtime - Lookup Object
foo := v1alpha1.Foo{}
err := c.Client().Get(context.TODO(), request.NamespacedName, foo)

////////////////////////////////////////////////////////////////////
// Owners References
////////////////////////////////////////////////////////////////////

// Sample Controller - Set OwnerRefrence
OwnerReferences: []metav1.OwnerReference{
	*metav1.NewControllerRef(foo, schema.GroupVersionKind{
	Group:   samplev1alpha1.SchemeGroupVersion.Group,
	Version: samplev1alpha1.SchemeGroupVersion.Version,
	Kind:    "Foo"})}

// Controller-Runtime - Set OwnerReference
controllerutil.SetControllerReference(foo, deployment, c.scheme)
