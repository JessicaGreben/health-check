package checks

import (
<<<<<<< HEAD
	"bytes"
	"fmt"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
)

// check best practice labels are defined on a deployment and pods.
func labels(deployment appsv1.Deployment) (bool, string) {
	violation := false
	var msg bytes.Buffer
	var lblExists bool

	collLabels := map[string]map[string]string{"deployment": deployment.GetObjectMeta().GetLabels(), "pod": deployment.Spec.Template.GetLabels()}
	deployment.Spec.Templ
	for _, lbl := range [1]string{"app"} {
		for lblFrom := range collLabels {
			_, lblExists = collLabels[lblFrom][lbl]
			if !lblExists {
				msg.WriteString(fmt.Sprintf("%s label '%s' does not exist\n", lblFrom, lbl))
				violation = true
			}
		}
	}
	return violation, msg.String()
}

// check for any pods in deployment spec using hostPort
func hostPort(deployment appsv1.Deployment) (bool, string) {
	for container := range deployment.Spec.Template.Spec.Containers {
		for port := range container.Ports {
			//TODO: check if HostPort is defined
		}
	}
}

// check if a struct field is omitted.  I have no idea if this will work.
func StructFieldExists(IFace interface{}, FieldName string) bool {
	ValueIface := reflect.ValueOf(Iface)

	//check if passed interface is a pointer
	if ValueIface.Type().Kind() != reflect.Ptr {
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	// deref with Elem() and get field by name
	Field := ValueIface.Elem().FieldByName(FieldName)
	if !Field.IsValid() {
		return false
	}
	return true
}
=======
	"github.com/jessicagreben/health-check/pkg/kube"
	"github.com/jessicagreben/health-check/pkg/types"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Deploys executes all the health checks for all pods.
func Deploys() ([]types.DeployResults, error) {
	var results []types.DeployResults

	deploys, err := kube.AppsV1API.Deployments("").List(metav1.ListOptions{})
	if err != nil {
		return results, err
	}

	for _, deploy := range deploys.Items {
		deployResults := types.DeployResults{
			Name:      deploy.Name,
			Namespace: deploy.Namespace,
		}

		for _, container := range deploy.Spec.Template.Spec.Containers {
			req, limits := resources(container)
			live, ready := probes(container)

			ctnResults := types.ContainerResults{
				Name:     container.Name,
				Requests: req,
				Limits:   limits,
				Live:     live,
				Ready:    ready,
			}
			deployResults.Containers = append(deployResults.Containers, ctnResults)
		}

		results = append(results, deployResults)
	}

	return results, nil
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

// avoid using "latest" for a tag for image
func containerTagCheck() {}

// TODO: add check that cpu limit is not > 1Gb
func resourceCPULimit() {}
>>>>>>> origin/master
