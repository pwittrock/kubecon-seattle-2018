type Controller struct {
	// directxman12: begin highlight
	sampleclientset clientset.Interface
	foosLister        listers.FooLister
	foosSynced        cache.InformerSynced
	// directxman12: end highlight
	workqueue workqueue.RateLimitingInterface
}

func Start() {
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	// directxman12: begin highlight
	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)
	fooInformer := exampleInformerFactory.Samplecontroller().V1alpha1().Foos()
	// directxman12: end highlight

	controller := &Controller{
		// directxman12: begin highlight
		sampleclientset: sampleclientset,
		foosLister:      fooInformer.Lister(),
		foosSynced:      fooInformer.Informer().HasSynced,
		// directxman12: end highlight

		workqueue:       workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Foos"),
	}

	// directxman12: begin highlight
	fooInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueFoo,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueFoo(new)
		},
	})

	exampleInformerFactory.Start(stopCh)
	// directxman12: end highlight

	controller.Run(2, stopCh)
}

// directxman12: begin highlight
func (c *Controller) enqueueFoo(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.AddRateLimited(key)
}
// directxman12: end highlight
