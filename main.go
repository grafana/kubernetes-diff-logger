package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/joe-elliott/kubernetes-diff-logger/pkg/differ"
	"github.com/joe-elliott/kubernetes-diff-logger/pkg/signals"
	"github.com/joe-elliott/kubernetes-diff-logger/pkg/wrapper"
)

var (
	masterURL    string
	kubeconfig   string
	resyncPeriod time.Duration
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.DurationVar(&resyncPeriod, "resync", time.Second*30, "Periodic interval in which to force resync objects.")
}

func main() {
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	// creates the clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("kubernetes.NewForConfig failed: %v", err)
	}

	informerFactory := informers.NewSharedInformerFactory(client, resyncPeriod)
	informer, wrap, err := informerForName("deployment", informerFactory)
	if err != nil {
		log.Fatalf("informerForName failed: %v", err)
	}

	differ.NewDiffer("", 30*time.Second, wrap, informer)

	stop := signals.SetupSignalHandler()
	informerFactory.Start(stop)
}

func informerForName(name string, i informers.SharedInformerFactory) (cache.SharedInformer, wrapper.Wrap, error) {

	switch name {
	case "deployment":
		return i.Apps().V1().Deployments().Informer(), wrapper.NewDeploymentWrapper, nil
	}

	return nil, nil, fmt.Errorf("Unsupported informer name %s", name)
}
