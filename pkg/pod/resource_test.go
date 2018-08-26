package pod

import (
	"testing"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPodsResourceCheck(t *testing.T) {
	poddie := v1.Pod{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.PodSpec{},
		Status:     v1.PodStatus{},
	}
	podz := make([]v1.Pod, 1)
	podzz := append(podz, poddie)

	cases := []struct {
		pods *v1.PodList
	}{
		{&v1.PodList{
			metav1.TypeMeta{},
			metav1.ListMeta{},
			podzz,
			},
		},
	}

	var emptyReport = PodReport{}

	for _, c := range cases {
		report := PodsResourceCheck(c.pods)
		if len(report.Limits) != 0 {
			t.Errorf("No pods: expected %v but returned %v", emptyReport, report)
		}
	}
}
