package sauce

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURL = "https://%s:%s@saucelabs.com/rest/v1/%s"
)

//Client representative of saucelabs client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//NewClient Generates new saucelabs client
func NewClient(apiKey, userName, baseURL string) *Client {

	if baseURL == "" {
		baseURL = fmt.Sprintf(BaseURL, userName, apiKey, userName)
	}

	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
