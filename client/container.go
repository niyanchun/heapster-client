package client

import "github.com/niyanchun/heapster-client/type/v1"

func (c *Client) ListContainer(namespace, podName string) ([]string, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/pods/" + podName + "/containers/"

	return c.list(url)
}

func (c *Client) ListContainerMetrics(namespace, podName, containerName string) ([]string, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/pods/" + podName + "/containers/" +
		containerName + "/metrics/"

	return c.list(url)
}

func (c *Client) GetContainerMetrics(namespace, podName, containerName, metricName, start, end string) (*v1.Metrics, error) {
	url := c.Url + "/api/v1/model/namespaces/" + namespace + "/pods/" + podName + "/containers/" +
		containerName + "/metrics/" + metricName

	url += handleQueryParams(start, end)

	return c.getMetrics(url)
}

func (c *Client) ListFreeContainer(nodeName string) ([]string, error) {
	url := c.Url + "/api/v1/model/nodes/" + nodeName + "/freecontainers/"

	return c.list(url)
}

func (c *Client) ListFreeContainerMetrics(nodeName, containerName string) ([]string, error) {
	url := c.Url + "/api/v1/model/nodes/" + nodeName + "/freecontainers/" + containerName + "/metrics/"

	return c.list(url)
}

func (c *Client) GetFreeContainerMetrics(nodeName, containerName, metricName, start, end string) (*v1.Metrics, error) {
	url := c.Url + "/api/v1/model/nodes/" + nodeName + "/freecontainers/" + containerName + "/metrics/" + metricName

	url += handleQueryParams(start, end)

	return c.getMetrics(url)
}
