package client

import "github.com/niyanchun/heapster-client/type/v1"

func (c *Client) ListNamespace() ([]string, error) {
	url := c.Url + "/api/v1/model/namespaces/"

	return c.list(url)
}

func (c *Client) ListNamespaceMetrics(namespace string) ([]string, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/metrics/"

	return c.list(url)
}

func (c *Client) GetNamespaceMetrics(namespace, metricName, start, end string) (*v1.Metrics, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/metrics/" + metricName

	url += handleQueryParams(start, end)

	return c.getMetrics(url)
}
