package client

import "testing"

func TestGetClusterMetrics(t *testing.T) {
	// point 1
	metrics, err := client.ListClusterMetric()
	if err != nil {
		t.Error(err)
	}

	if len(metrics) == 0 {
		t.Error("No available cluster metrics")
	}

	t.Logf("ListClusterMetric:\n: %v", metrics)

	// point 2
	t.Log("GetClusterMetrics")
	start, end := getTimeRange()
	for _, metric := range metrics {

		result, err := client.GetClusterMetrics(metric, start, end)
		if err != nil {
			t.Error(err)
		}

		if result == nil {
			t.Errorf("metric(%s)'s result is empty", metric)
		}

		t.Logf("%s:\n %v", metric, result)
	}
}
