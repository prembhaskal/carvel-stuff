package pkglist

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

	var unstructured unstructured.Unstructured

	// get package(s)
	err = restClient.Get().
	// Namespace("default").
	// Resource("configmaps").
	Namespace("sleeptest").
	Resource("packages").
	Timeout(timeout).
	Do(context.TODO()).
	Into(&unstructured)
	// print data 
	if err != nil {
		log.Fatalf("error pkg list: %v, err-type: %t", err, err)
	}

	fmt.Printf("got result: %v\n", unstructured)
	
	// done
}

func setConfigDefaults(config *rest.Config) {
	gv := schema.GroupVersion{
		Group: "data.packaging.carvel.dev",
		Version: "v1alpha1",
	}
	// gv := schema.GroupVersion{
	// 	Group: "",
	// 	Version: "v1",
	// }
	config.GroupVersion = &gv
	config.APIPath = "/apis" // not "api", do kubectl -n sleeptest get pkg --v=6 to get the API call made there.

	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
}