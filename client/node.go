package client

import "github.com/niyanchun/heapster-client/type/v1"

func (c *Client) ListNodes() ([]string, error) {
	url := c.Url + "/api/v1/model/nodes/"

	return c.list(url)
}

func (c *Client) ListNodeMetric(nodeName string) ([]string, error) {
	url := c.Url + "/api/v1/model/nodes/" + nodeName + "/metrics/"

	return c.list(url)
}

func (c *Client) GetNodeMetrics(nodeName, metricName, start, end string) (*v1.Metrics, error) {
	url := c.Url + "/api/v1/model/nodes/" + nodeName + "/metrics/" + metricName

	url += handleQueryParams(start, end)

	return c.getMetrics(url)
}
