type Controller struct {}

func Start() {
	ctrl, _ := builder.SimpleController().Build(&Controller{})
	ctrl.Start(signals.SetupSignalHandler())
}
