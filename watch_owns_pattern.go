type Controller struct {
	Client client.Client
}

func (ct *Controller) InjectClient(c client.Client) error {
	ct.Client = c
	return nil
}

func Start() {
	ctrl, _ := builder.SimpleController().
		ForType(&v1alpha1.Foo{}).
		// directxman12: begin highlight
		Owns(&appsv1.Deploment{}).
		// directxman12: end highlight
		Build(&Controller{})
	ctrl.Start(signals.SetupSignalHandler())
}

