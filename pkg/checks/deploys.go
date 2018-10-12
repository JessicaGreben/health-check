package checks

import (
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
