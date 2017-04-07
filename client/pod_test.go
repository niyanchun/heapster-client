package client

import "testing"

const SYS_NAMESPACE string = "kube-system"

func TestPod(t *testing.T) {
	// point 1
	pods, err := client.ListPod(SYS_NAMESPACE)
	if err != nil {
		t.Error(err)
	}

	if len(pods) == 0 {
		t.Error("No pods in namespace kube-system")
	}
	t.Logf("ListPod:\n%v", pods)

	// point 2
	metrics, err := client.ListPodMetrics(SYS_NAMESPACE, pods[0])
	if err != nil {
		t.Error(err)
	}

	if len(metrics) == 0 {
		t.Error("No available pod metric")
	}
	t.Logf("ListPodMetrics:\n%v", metrics)

	// point 3
	t.Log("GetPodMetrics:")
	start, end := getTimeRange()
	for _, metric := range metrics {
		result, err := client.GetPodMetrics(SYS_NAMESPACE, pods[0], metric, start, end)
		if err != nil {
			t.Error(err)
		}

		if result == nil {
			t.Errorf("Pod metric(%s) result is empty", metric)
		}

		t.Logf("%s:\n %v", metric, result)
	}

}
