package codeclimateclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiKey  string
	BaseUrl string
}

// TODO: Extend in the future to accept POST requests
func (c *Client) makeRequest(method string, path string, payload io.Reader) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.BaseUrl, path), payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.ApiKey))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}
