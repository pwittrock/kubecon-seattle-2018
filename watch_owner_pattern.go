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

func Start() {
	// Watch ReplicaSet Changes and Trigger MyController Reconcile Function
	ctrl, _ := builder.SimpleController().
		ForType(&v1alpha1.Foo{}).
		Owns(&appsv1.Deploment{}).
		Build(&Controller{})
	ctrl.Start(signals.SetupSignalHandler())
}

type Controller struct {
	Client client.Client
}

func (ct *Controller) InjectClient(c client.Client) error {
	ct.Client = c
	return nil
}