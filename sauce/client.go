package sauce

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// sauce returns responses as plain text, account for messages here
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
		msg, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}

		errRes := errorResponse{
			Code:    res.StatusCode,
			Message: string(msg),
		}

		return errors.New(errRes.Message)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
