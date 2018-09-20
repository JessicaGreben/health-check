package kubeconf

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

// CoreV1API exports the CoreV1 API client.
var CoreV1API = clientset.CoreV1()

// AutoscalingV2beta1API exports the AutoscalingAPI client.
var AutoscalingV2beta1API = clientset.AutoscalingV2beta1()
