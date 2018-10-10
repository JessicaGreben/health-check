package kube

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// createClientset loads kubeconfig and setups the connection to the k8s api.
func createClientset() *kubernetes.Clientset {
	kubeconfig := filepath.Join(
		os.Getenv("HOME"),
		".kube",
		"config",
	)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("Error:", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return clientset
}

var clientset = createClientset()

// Kubernetes version 1.10 APIs

// AutoscalingV2beta1API exports the AutoscalingAPI client.
var AutoscalingV2beta1API = clientset.AutoscalingV2beta1()

// Kubernetes version 1.11 APIs

// CoreV1API exports the v1 Core API client.
var CoreV1API = clientset.CoreV1()

// AutoscalingV1API exports the v1 Autoscaling API client.
var AutoscalingV1API = clientset.AutoscalingV1()

// AppsV1API exports the v1 Apps API client.
var AppsV1API = clientset.AppsV1()
