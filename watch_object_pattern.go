type Controller struct {
	// directxman12: begin highlight
	Client client.Client
	// directxman12: end highlight
}

// directxman12: begin highlight
func (ct *Controller) InjectClient(c client.Client) error {
	ct.Client = c
	return nil
}
// directxman12: end highlight

func Start() {
	ctrl, _ := builder.SimpleController().
		// directxman12: begin highlight
		ForType(&v1alpha1.Foo{}).
		// directxman12: end highlight
		Build(&Controller{})
	ctrl.Start(signals.SetupSignalHandler())
}
