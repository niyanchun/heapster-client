package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/niyanchun/heapster-client/type/v1"
)

type Client struct {
	HttpClient *http.Client
	Url        string
	Username   string
	Password   string
}

// NewClient creates a new client.
func NewClient(url, username, password string) Client {
	httpClient := &http.Client{}

	return Client{HttpClient: httpClient, Url: url, Username: username, Password: password}
}

func (c *Client) list(url string) ([]string, error) {
	var result []string
	err := c.get(url, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) getMetrics(url string) (*v1.Metrics, error) {
	var result v1.Metrics
	err := c.get(url, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c *Client) createRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(c.Username, c.Password)

	return request, nil
}

func (c *Client) get(url string, result interface{}) error {
	log.Printf("Requesting GET %v\n", url)

	req, err := c.createRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		printResponse(resp)

		return errors.New("Http request return status code: " + resp.Status)
	}

	//printResponse(resp)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func printResponse(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	log.Printf("%s\n", string(body))
}

func handleQueryParams(start, end string) string {
	var query string

	if len(start) != 0 {
		query = "?start=" + start
	} else {
		query += "?start=" + "1970-00-00T00:00:00Z"
	}

	if len(end) != 0 {
		query += "&end=" + end
	}
	return query
}
