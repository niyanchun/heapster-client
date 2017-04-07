# Heapster-client

## Introduction

[Heapster](https://github.com/kubernetes/heapster) enables Container Cluster Monitoring and Performance Analysis. And it has some APIs.
This project is a library that makes it easier for you to use the APIs.

NB: If you are not familiar to the Heapster model, I suggest you read [Heapster Metric Model](https://github.com/kubernetes/heapster/blob/master/docs/model.md) first.

## Usage

1. Get the library

    ```
    go get github.com/niyanchun/heapster-client
    ```

2. List Cluster-Level metric, and get `cpu/usage_rate` metric statistics:

    ```
    package main

    import (
        "fmt"
        "github.com/niyanchun/heapster-client/client"
        "time"
        "github.com/niyanchun/heapster-client/type/v1"
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

    ```
    
## Compatibility

This library is test under Kubernetes 1.6.0 and Heapster v1.3.0-beta.0.