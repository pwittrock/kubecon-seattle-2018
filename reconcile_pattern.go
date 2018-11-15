////////////////////////////////////////////////////////////////////
// Reconcile Request - Get Object Key
////////////////////////////////////////////////////////////////////

// Sample Controller
namespace, name, err := cache.SplitMetaNamespaceKey(key)
if err != nil {
	runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
	return nil
}

// Controller-Runtime
namespace := request.Namespace
name := request.Name

////////////////////////////////////////////////////////////////////
// Client - Lookup Object
////////////////////////////////////////////////////////////////////

// Sample Controller
foo, err := c.foosLister.Foos(namespace).Get(name)

// Controller-Runtime
foo := v1alpha1.Foo{}
err := c.Client().Get(context.TODO(), request.NamespacedName, foo)

////////////////////////////////////////////////////////////////////
// Owners References - Set OwnerRefrence
////////////////////////////////////////////////////////////////////

// Sample Controller
OwnerReferences: []metav1.OwnerReference{
	*metav1.NewControllerRef(foo, schema.GroupVersionKind{
	Group:   samplev1alpha1.SchemeGroupVersion.Group,
	Version: samplev1alpha1.SchemeGroupVersion.Version,
	Kind:    "Foo"})}

// Controller-Runtime
controllerutil.SetControllerReference(foo, deployment, c.scheme)
