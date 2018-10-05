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
	kubeconfig := getKubeConfig()

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

// get a valid kubeconfig path
func getKubeConfig() string {
	var kPath string
	if os.Getenv("KUBECONFIG") != "" {
		kPath = os.Getenv("KUBECONFIG")
	} else if home := os.Getenv("HOME"); home != "" {
		kPath = filepath.Join(home, ".kube", "config")
	} else {
		fmt.Println("kubeconfig not found.  Please ensure ~/.kube/config exists or KUBECONFIG is set.")
		os.Exit(1)
	}

	if _, err := os.Stat(kPath); err != nil {
		//kubeconfig doesn't exist
		fmt.Printf("%s doesn't exist - do you have a kubeconfig configured?\n", kPath)
		os.Exit(1)
  }
	return kPath
}

var clientset = createClientset()

// CoreV1API exports the CoreV1 API client.
var CoreV1API = clientset.CoreV1()

// AutoscalingV2beta1API exports the AutoscalingAPI client.
var AutoscalingV2beta1API = clientset.AutoscalingV2beta1()
