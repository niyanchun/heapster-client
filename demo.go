package main

import (
	"fmt"
	"github.com/niyanchun/heapster-client/client"
	"github.com/niyanchun/heapster-client/type/v1"
	"time"
)

func main() {
	HEAPSTER_URL := "http://192.168.56.101:8080/api/v1/proxy/namespaces/kube-system/services/heapster"
	client := client.NewClient(HEAPSTER_URL, "", "")

	start, end := getTimeRange()
	list, err := client.ListClusterMetric()
	checkErr(err)
	fmt.Printf("ListClusterMetric:\n %v", list)

	metrics, err := client.GetClusterMetrics(v1.CPU_USAGE_RATE, start, end)
	checkErr(err)
	fmt.Printf("ListClusterMetric:\n %v", metrics)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getTimeRange() (string, string) {
	now := time.Now().UTC()
	end := now.Format(time.RFC3339)
	start := now.Add(-600 * 1e9).Format(time.RFC3339) // ten minutes before
	fmt.Println(start, end)

	return start, end
}
