package kube

import (
	"github.com/aSquidsBody/go-common/logs"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Config() *rest.Config {
	c, err := rest.InClusterConfig()
	if err != nil {
		logs.Fatal("Could not create InClusterConfig", err)
	}

	return c
}

func NewClientset() *kubernetes.Clientset {
	c := Config()
	clientSet, err := kubernetes.NewForConfig(c)
	if err != nil {
		logs.Fatal("Could not create Clientset", err)
	}
	return clientSet
}
