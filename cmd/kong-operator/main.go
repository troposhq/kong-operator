package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/troposhq/kong-operator/internal/pkg/controller"
	"github.com/troposhq/kong-operator/pkg/signals"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	controller := controller.New(kubeClient)
	if err = controller.Run(2, stopCh); err != nil {
		log.Fatalf("Error running controller: %s", err.Error())
	}
}
