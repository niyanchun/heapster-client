package client

import "testing"

func TestNode(t *testing.T) {
	// point 1
	nodes, err := client.ListNodes()
	if err != nil {
		t.Error(err)
	}

	if len(nodes) == 0 {
		t.Error("No available Node")
	}

	t.Logf("ListNodes:\n%v", nodes)

	// point 2
	metrics, err := client.ListNodeMetric(nodes[0])
	if err != nil {
		t.Error(err)
	}

	if len(metrics) == 0 {
		t.Errorf("No metric for node: %s", nodes[0])
	}

	t.Logf("ListNodeMetric:\n%v", metrics)

	// point 3
	t.Log("GetNodeMetrics:")
	start, end := getTimeRange()
	for _, metric := range metrics {
		result, err := client.GetNodeMetrics(nodes[0], metric, start, end)
		if err != nil {
			t.Error(err)
		}

		if result == nil {
			t.Errorf("node's metric(%s) result is empty", metric)
		}

		t.Logf("%s:\n %v", metric, result)
	}
}
