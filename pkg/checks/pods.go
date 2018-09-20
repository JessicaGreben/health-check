package checks

import (
	"fmt"

	"github.com/jessicagreben/pratique/pkg/kubeconf"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RunPodHealthChecks runs all the health checks for a list of pods
func RunPodHealthChecks() {
	coreV1API := kubeconf.CoreV1API
	pods, err := coreV1API.Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, pod := range pods.Items {
		fmt.Printf("\nRunning health check for pod: %s", pod.Name)
		fmt.Printf("\nNamespace: %s", pod.Namespace)
		fmt.Printf("\nThis pod has %d container/s.", len(pod.Spec.Containers))
		fmt.Print("\n==============================================\n")
		runHPACheck(pod.Name, pod.Namespace)
		nakedPodCheck(pod)

		for _, container := range pod.Spec.Containers {
			fmt.Printf("Checking pod container: %s\n", container.Name)
			runResourceCheck(container)
			runProbeCheck(container)
		}
	}
}

// do not use naked pods, they will not be rescheduled in the event of failure
func nakedPodCheck(pod v1.Pod) {
	fmt.Print("Naked Pod check:\n")
	podmeta := pod.GetObjectMeta()
	ref := podmeta.GetOwnerReferences()
	if len(ref) == 0 {
		fmt.Print("- failed. This pod is naked and is not bound to an ReplicaSet or Deplpyment.\n")
		return
	}
	fmt.Print("- pass.")
}

// are pod resources using the latest stable API version available
func latestVersionCheck() {
	// get version of cluster
	// find out the latest version of the API
	// check if current resource is that version
}

// Don’t specify a hostPort for a Pod unless it is absolutely necessary
// ref: https://kubernetes.io/docs/concepts/configuration/overview/
func bindHostPort() {}

// avoid using latest tag for containers
func containerTagCheck() {}

func resourceCPULimit() {}

// PDB view shows how much "can" be disrupted - so if `ALLOWED DISRUPTIONS < 1`
// then a deadlock will definitely occur on upgrades
func pdbCheck() {}

// how many loadbalancers if cluster
func loadBalancerCheck() {}

func runResourceCheck(container v1.Container) {
	fmt.Print("Resource check:\n")
	// TODO: add check that cpu limit is not > 1Gb
	// if it is, recommend reducing the cpu limit and
	// instead increase the replica count
	// this can help with scheduling
	// ref: https://youtu.be/xjpHggHKm78?t=1m55s
	// resourceCPULimit()
	if len(container.Resources.Limits) == 0 {
		fmt.Printf("- No resource limits set.\n")
	}
	if len(container.Resources.Requests) == 0 {
		fmt.Printf("- No resource requests set.\n")
	}
	// TODO: add recommendation to possibly use a LimitRange
	// resource to set defaults, max, and min values
	// ref: https://kubernetes.io/docs/tasks/administer-cluster/manage-resources/cpu-default-namespace/
	if len(container.Resources.Limits) != 0 && len(container.Resources.Requests) != 0 {
		fmt.Print("- Pass.\n")
	}
}

func runProbeCheck(container v1.Container) {
	fmt.Print("Probe check:\n")
	if container.LivenessProbe == nil {
		fmt.Printf("- No liveness probe\n")
	}
	if container.ReadinessProbe == nil {
		fmt.Printf("- No readiness probe set.\n")
	}
	if container.ReadinessProbe != nil && container.LivenessProbe != nil {
		fmt.Print("- Pass.\n")
	}
}

func runHPACheck(podName string, namespace string) {
	fmt.Print("HPA check:\n")
	autoscaleAPI := kubeconf.AutoscalingV2beta1API
	hpas, err := autoscaleAPI.HorizontalPodAutoscalers(namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if len(hpas.Items) == 0 {
		fmt.Print("- No HPA configured.\n")
		return
	}
	fmt.Print("The following HPAs exist for this pod:\n")
	for _, item := range hpas.Items {
		name := item.GetObjectMeta().GetName()
		fmt.Print("- ", name, "\n")
	}
}
