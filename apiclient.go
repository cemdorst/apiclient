package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client -
type Client struct {
	ConfigHostURL string
	ConfigToken   string
	HTTPClient    *http.Client
	Payload       string
}

// NewClient -
func (c *Client) New(token, apiAddress string) {

	c.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	c.ConfigHostURL = apiAddress
	c.ConfigToken = "Token " + token
}

// Do Generic Request
func (c *Client) DoRequest(method, url string) ([]byte, error) {

	var idns map[string]interface{}
	var result []byte

	payload := strings.NewReader(c.Payload)

	req, err := http.NewRequest(method, c.ConfigHostURL+url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.ConfigToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json; version=3")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	switch code := resp.StatusCode; {
	case code != 200:
		err := fmt.Errorf("%d", resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()
	result, _ = ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(result, &idns); err != nil {
		return nil, err
	}

	return result, nil
}
