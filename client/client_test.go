package client

import (
	"os"
	"testing"
	"time"
)

var HEAPSTER_URL string

var client Client

func TestMain(m *testing.M) {
	HEAPSTER_URL = "http://192.168.56.101:8080/api/v1/proxy/namespaces/kube-system/services/heapster"

	client = NewClient(HEAPSTER_URL, "", "")
	os.Exit(m.Run())
}

func getTimeRange() (string, string) {
	now := time.Now().UTC()
	end := now.Format(time.RFC3339)
	start := now.Add(-600 * 1e9).Format(time.RFC3339) // ten minutes before

	return start, end
}
