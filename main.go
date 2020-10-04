package main

import (
	"errors"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	client, err := newInClusterClient()
	if err != nil {
		panic(err.Error())
	}

	path := "/healthz"
	content, err := client.Discovery().RESTClient().Get().AbsPath(path).DoRaw()
	if err != nil {
		fmt.Printf("ErrorBadRequst : %s\n", err.Error())
		os.Exit(1)
	}

	contentStr := string(content)
	if contentStr != "ok" {
		fmt.Printf("ErrorNotOk : response != 'ok' : %s\n", contentStr)
		os.Exit(1)
	}

	fmt.Printf("Success : ok\n")
	os.Exit(0)
}

func newInClusterClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return &kubernetes.Clientset{}, errors.New("Failed loading client config")
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &kubernetes.Clientset{}, errors.New("Failed getting clientset")
	}
	return clientset, nil
}
