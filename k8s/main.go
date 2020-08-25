package main

import (
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

func main() {
	kubeconfig := filepath.Join(
		"~/Downloads/shantanu-spectro-app-aws", ".kube", "config",
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

}
