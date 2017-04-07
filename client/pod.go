package client

import "github.com/niyanchun/heapster-client/type/v1"

func (c *Client) ListPod(namespace string) ([]string, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/pods/"

	return c.list(url)
}

func (c *Client) ListPodMetrics(namespace, podName string) ([]string, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/pods/" + podName + "/metrics/"

	return c.list(url)
}

func (c *Client) GetPodMetrics(namespace, podName, metricName, start, end string) (*v1.Metrics, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/pods/" + podName + "/metrics/" + metricName

	url += handleQueryParams(start, end)

	return c.getMetrics(url)
}
