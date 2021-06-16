package sauce

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	var err error

	if body == nil {
		req, err = http.NewRequest(method, endpoint, nil)
	} else {
		req, err = http.NewRequest(method, endpoint, body)
	}

	//check if request built without error
	if err != nil {
		log.Fatal(err)
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
		msg, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("cant read body")
		}

		return errors.New(string(msg))
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
