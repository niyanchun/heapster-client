package client

import "testing"

func TestNamespace(t *testing.T) {
	// point 1
	ns, err := client.ListNamespace()
	if err != nil {
		t.Error(err)
	}

	if len(ns) == 0 {
		t.Error("result is empty")
	}

	t.Logf("ListNamespace:\n%v", ns)

	// point 2
	metrics, err := client.ListNamespaceMetrics(ns[0])
	if err != nil {
		t.Error(err)
	}

	if len(metrics) == 0 {
		t.Error("No available namesapce metrics")
	}

	t.Logf("ListNamespaceMetrics:\n%v", metrics)

	// point 3
	t.Log("GetNamespaceMetrics:")
	start, end := getTimeRange()
	for _, metric := range metrics {
		result, err := client.GetNamespaceMetrics(ns[0], metric, start, end)
		if err != nil {
			t.Error(err)
		}

		if result == nil {
			t.Errorf("namespace's metric(%s) result is empty.", metric)
		}

		t.Logf("%s:\n %v", metric, result)
	}
}
