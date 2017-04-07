package client

import "github.com/niyanchun/heapster-client/type/v1"

// List Cluster-Level Metric
func (c *Client) ListClusterMetric() ([]string, error) {
	url := c.Url + "/api/v1/model/metrics/"

	return c.list(url)
}

func (c *Client) GetClusterMetrics(metricName, start, end string) (*v1.Metrics, error) {
	url := c.Url + "/api/v1/model/metrics/" + metricName

	url += handleQueryParams(start, end)

	return c.getMetrics(url)
}
