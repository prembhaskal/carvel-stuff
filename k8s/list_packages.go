package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes/scheme"
)

func ListPackages() {
	// config
	kubeConfigEnv := os.Getenv("KUBECONFIG")
	if kubeConfigEnv == "" {
		log.Fatalf("no kubeconfig defined")
	}
	kubeBytes, err := os.ReadFile(kubeConfigEnv)
	if err != nil {
		log.Fatalf("error reading kubeconfig at %s: %v", kubeConfigEnv, kubeBytes)
	}

	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeBytes)
	if err != nil {
		log.Fatalf("error creating REST config: %v", err)
	}

	setConfigDefaults(restConfig)

	// create client
	restClient, err := rest.RESTClientFor(restConfig)
	if err != nil {
		log.Fatalf("error client create: %v", err)
	}

	timeout := time.Second * 30


	// get package(s)
	result := restClient.Get().
	Namespace("sleeptest").
	Resource("packages").
	Timeout(timeout).
	Do(context.TODO())
	// print data 

	fmt.Printf("got result: %v\n", result)
	
	// done
}

func setConfigDefaults(config *rest.Config) {
	gv := schema.GroupVersion{
		Group: "data.packaging.carvel.dev",
		Version: "v1alpha1",
	}
	config.GroupVersion = &gv
	config.APIPath = "/api"

	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
}