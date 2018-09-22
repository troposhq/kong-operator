package controller

import (
	"k8s.io/client-go/kubernetes"
)

type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
}

func New(kubeclientset kubernetes.Interface) *Controller {
	controller := &Controller{
		kubeclientset: kubeclientset,
	}
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
