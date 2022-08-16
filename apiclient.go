package apiclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client -
type Client struct {
	ConfigHostURL string
	ConfigToken   string
	HTTPClient    *http.Client
}

// NewClient -
func (c *Client) New(token, apiAddress string) {

	c.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	c.ConfigHostURL = apiAddress
	c.ConfigToken = "Token " + token
}

//Do Generic Request
func (c *Client) DoRequest(method, url string) (result []byte) {

	var idns map[string]interface{}

	req, err := http.NewRequest(method, c.ConfigHostURL+url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", c.ConfigToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json; version=3")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, _ = ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(result, &idns); err != nil {
		fmt.Println(err)
	}

	for k, j := range idns {
		fmt.Printf("%s, type: %T, %s, type: %T \n", k, k, j, j)
	}

	return
}
