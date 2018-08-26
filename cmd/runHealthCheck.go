package cmd

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"k8s.io/api/core/v1"
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

type PodReport struct {
	limits    []string
	requests  []string
	liveness  []string
	readiness []string
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
	podReports := podsResourceCheck(pods)
	fmt.Println("pods with no limits:", podReports.limits)
	fmt.Println("pods with no request:", podReports.requests)
	fmt.Println("pods with no liveness:", podReports.liveness)
	fmt.Println("pods with no readiness:", podReports.readiness)
}

// Return a report that lists pods with no resource requests/limits
func podsResourceCheck(pods *v1.PodList) PodReport {
	var noRequests []string
	var noLimits []string
	var noLiveness []string
	var noReadiness []string
	fmt.Print("\nPod Resource Check:\n")
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			if len(container.Resources.Limits) == 0 {
				noLimits = append(noLimits, pod.Name)
			}
			if len(container.Resources.Requests) == 0 {
				noRequests = append(noRequests, pod.Name)
			}
			if container.LivenessProbe == nil {
				noLiveness = append(noLiveness, pod.Name)
			}
			if container.ReadinessProbe == nil {
				noReadiness = append(noReadiness, pod.Name)
			}
		}
	}
	report := PodReport{limits:noLimits, requests:noRequests,liveness:noLiveness, readiness:noReadiness}
	return report
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
