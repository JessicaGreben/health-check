package cmd

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/jessicagreben/pratique/pkg/pod"

	// "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// runHealthCheckCmd represents the runHealthCheck command
var runHealthCheckCmd = &cobra.Command{
	Use:   "runHealthCheck",
	Short: "Check the health of the current Kubernetes cluster",
	Long: `Run a health check against a Kubernetes cluster. A report is generated describing the results of the checks.`,
	Run: runHealthCheck,
}

func runHealthCheck(cmd *cobra.Command, args []string) {
	fmt.Println("runHealthCheck called")

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

	// You can see the available clientsets:
	// https://github.com/kubernetes/client-go/blob/master/kubernetes/clientset.go
	coreV1Api := clientset.CoreV1()

	pods, err := coreV1Api.Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error:", err)
	}
	podReports := pod.PodsResourceCheck(pods)
	fmt.Println("pods with no limits:", podReports.Limits)
	fmt.Println("pods with no request:", podReports.Requests)
	fmt.Println("pods with no liveness:", podReports.Liveness)
	fmt.Println("pods with no readiness:", podReports.Readiness)
}

func init() {
	rootCmd.AddCommand(runHealthCheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runHealthCheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runHealthCheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
