package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default API URL
const HostURL string = "https://api.azionapi.net:443"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default API URL
		HostURL: HostURL,
		Token:   *token,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
