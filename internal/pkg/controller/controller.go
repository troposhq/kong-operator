package controller

type Controller struct {
}

func NewController() *Controller {
	controller := &Controller{}
	return controller
}

func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	return true
}
