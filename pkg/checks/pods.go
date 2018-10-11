package checks

import (
	//"github.com/jessicagreben/health-check/pkg/kubeconf"
	"../kubeconf"

	"github.com/jessicagreben/health-check/pkg/types"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pods executes all the health checks for all pods.
func Pods() (types.Results, error) {
	var results types.Results

	pods, err := kubeconf.CoreV1API.Pods("").List(metav1.ListOptions{})
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

		for _, container := range pod.Spec.Containers {
			req, limits := resources(container)
			live, ready := probes(container)

			ctrResults := types.ContainerResults{
				Name:     container.Name,
				Requests: req,
				Limits:   limits,
				Live:     live,
				Ready:    ready,
			}
			podResults.Containers = append(podResults.Containers, ctrResults)
		}
		results.Pods = append(results.Pods, podResults)
	}

	return results, nil
}

// naked checks if there are any naked pods (i.e. pods without a
// replicaset).
// Naked pods should be avoided since they will not be rescheduled in the
// event of failure
func naked(pod v1.Pod) bool {
	podmeta := pod.GetObjectMeta()
	ref := podmeta.GetOwnerReferences()
	if len(ref) == 0 {
		return true
	}
	return false
}

// resources checks if there are resource requests and resource limits set.
func resources(container v1.Container) (bool, bool) {
	req, limits := true, true
	if len(container.Resources.Requests) == 0 {
		req = false
	}
	if len(container.Resources.Limits) == 0 {
		limits = false
	}
	return req, limits
}

// probes checks if there are liveness and readiness probes set.
func probes(container v1.Container) (bool, bool) {
	live, ready := true, true
	if container.LivenessProbe == nil {
		live = false
	}
	if container.ReadinessProbe == nil {
		ready = false
	}
	return live, ready
}

// HPAs checks if there are HPAs configured for a pod.
func HPAs(podName string, namespace string) (bool, error) {
	autoscaleAPI := kubeconf.AutoscalingV2beta1API
	hpas, err := autoscaleAPI.HorizontalPodAutoscalers(namespace).List(metav1.ListOptions{})
	if err != nil {
		return false, err
	}
	if len(hpas.Items) == 0 {
		return false, nil
	}
	return true, nil
}

// Labels checks for best practice labels are configured for a pod.
func labels(container v1.Container) (bool, bool) {

}
