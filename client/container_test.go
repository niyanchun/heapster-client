package client

import "testing"

func TestContainer(t *testing.T) {
	// point 1
	pods, err := client.ListPod(SYS_NAMESPACE)
	if err != nil {
		t.Error(err)
	}

	if len(pods) == 0 {
		t.Error("No pods in namespace kube-system")
	}

	containers, err := client.ListContainer(SYS_NAMESPACE, pods[0])
	if err != nil {
		t.Error(err)
	}

	if len(containers) == 0 {
		t.Error("No available containers")
	}
	t.Logf("ListContainer:\n%v", containers)

	// point 2
	metrics, err := client.ListContainerMetrics(SYS_NAMESPACE, pods[0], containers[0])
	if err != nil {
		t.Error(err)
	}

	if len(metrics) == 0 {
		t.Error("No available container metrics")
	}

	t.Logf("ListContainerMetrics metrics:\n%v", metrics)

	// point 3
	start, end := getTimeRange()
	for _, metric := range metrics {
		result, err := client.GetContainerMetrics(SYS_NAMESPACE, pods[0], containers[0], metric, start, end)
		if err != nil {
			t.Error(err)
		}

		if result == nil {
			t.Errorf("Container's metric(%s) result is empty", metric)
		}

		t.Logf("%s:\n %v", metric, result)
	}
}

/*
func TestFreeContainer(t *testing.T) {
	// point 1
	nodes, err := client.ListNodes()
	if err != nil {
		t.Error(err)
	}

	if len(nodes) == 0 {
		t.Error("No available Node")
	}

	containers, err := client.ListFreeContainer(nodes[0])
	if err != nil {
		t.Error(err)
	}

	if len(containers) == 0 {
		t.Error("No available free containers")
	}

	t.Logf("free containers: %v", containers)

	// point 2
	metrics, err := client.ListFreeContainerMetrics(nodes[0], containers[0])
	if err != nil {
		t.Error(err)
	}

	if len(metrics) == 0 {
		t.Error("No available free container metrics")
	}

	// point 3
	start, end := getTimeRange()
	for _, metric := range metrics {
		result, err := client.GetFreeContainerMetrics(nodes[0], containers[0], metric, start, end)
		if err != nil {
			t.Error(err)
		}

		if result == nil {
			t.Errorf("Free container's metric(%s) result is empty", metric)
		}
	}
}
*/
