package pod

import (
  "fmt"
  "k8s.io/api/core/v1"
)

type PodReport struct {
	Limits    []string
	Requests  []string
	Liveness  []string
	Readiness []string
}

// Return a report that lists pods with no resource requests/limits
func PodsResourceCheck(pods *v1.PodList) PodReport {
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
	report := PodReport{Limits:noLimits, Requests:noRequests,Liveness:noLiveness, Readiness:noReadiness}
	return report
}
