package sauce

import (
	"net/http"
	"testing"
)

func TestClientCreatedDefaultURL(t *testing.T) {
	c := NewClient("test", "test", "")
	if c != nil {
		t.Log("Client created successfully")

		if c.BaseURL == "https://test:test@saucelabs.com/rest/v1/test" {
			t.Log("BaseURL Generated successfully")
		}
	}
}

func TestClientBadRequest(t *testing.T) {
	c := NewClient("test", "test", "")
	err := c.sendRequest(http.NewRequest("POST", "", nil))
	if err != nil {
		t.Log("Sent bad request")
	}
}
