package checks

import (
	"github.com/jessicagreben/health-check/pkg/kube"
	"github.com/jessicagreben/health-check/pkg/types"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pods executes all the health checks for all pods.
func Pods() ([]types.PodResults, error) {
	var results []types.PodResults

	pods, err := kube.CoreV1API.Pods("").List(metav1.ListOptions{})
	if err != nil {
		return results, err
	}

	for _, pod := range pods.Items {
		hpa, err := HPAs(pod.Name, pod.Namespace)
		if err != nil {
			return results, err
		}

		podResults := types.PodResults{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			HPA:       hpa,
			Naked:     naked(pod),
		}

		results = append(results, podResults)
	}

	return results, nil
}

// naked checks if there are any naked pods (i.e. pods without a
// replicaset).
func naked(pod v1.Pod) bool {
	podmeta := pod.GetObjectMeta()
	ref := podmeta.GetOwnerReferences()
	if len(ref) == 0 {
		return true
	}
	return false
}

// HPAs checks if there are HPAs configured for a pod.
func HPAs(podName string, namespace string) (bool, error) {
	autoscaleAPI := kube.AutoscalingV1API
	hpas, err := autoscaleAPI.HorizontalPodAutoscalers(namespace).List(metav1.ListOptions{})
	if err != nil {
		return false, err
	}
	if len(hpas.Items) == 0 {
		return false, nil
	}
	return true, nil
}
