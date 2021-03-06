package sauce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://%s:%s@saucelabs.com/rest/v1/%s"
)

//Client representative of saucelabs client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

//NewClient Generates new saucelabs client
func NewClient(apiKey, userName, url string) *Client {

	if url == "" {
		url = fmt.Sprintf(baseURL, userName, apiKey, userName)
	}

	return &Client{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) buildRequest(method, endpoint string, body io.Reader) *http.Request {
	var req *http.Request

	if body == nil {
		req, _ = http.NewRequest(method, endpoint, nil)
	} else {
		req, _ = http.NewRequest(method, endpoint, body)
	}

	return req
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: " + fmt.Sprint(res.StatusCode))
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
