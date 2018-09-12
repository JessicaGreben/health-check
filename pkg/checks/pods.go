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
		fmt.Printf("\nIn %s namespace.", pod.Namespace)
		fmt.Printf("\nThis pod has %d container/s.", len(pod.Spec.Containers))
		fmt.Print("\n==============================================\n")
		runHPACheck(pod.Name, pod.Namespace)
		for _, container := range pod.Spec.Containers {
			fmt.Printf("Checking pod container: %s\n", container.Name)
			runResourceCheck(container)
			runProbeCheck(container)
		}
	}
}

func runResourceCheck(container v1.Container) {
	fmt.Print("Resource check:\n")
	if len(container.Resources.Limits) == 0 {
		fmt.Printf("- No resource limits set.\n")
	}
	if len(container.Resources.Requests) == 0 {
		fmt.Printf("- No resource requests set.\n")
	}
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
	autoscaleAPI := kubeconf.AutoscalingAPI
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
