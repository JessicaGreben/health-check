package checks

import (
	"bytes"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
)

// check best practice labels are defined on a deployment and pods.
func labels(deployment appsv1.Deployment) (bool, string) {
	violation := false
	var msg bytes.Buffer
	var lblExists bool

	collLabels := map[string]map[string]string{"deployment": deployment.GetObjectMeta().GetLabels(), "pod": deployment.Spec.Template.GetLabels()}

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
