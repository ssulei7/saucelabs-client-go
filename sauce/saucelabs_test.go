package sauce

import (
	"fmt"
	"net/http"
	"strings"
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

func TestCreateRequestValid(t *testing.T) {
	c := NewClient("test", "test", "")

	//create a valid request
	req := c.buildRequest("GET", fmt.Sprintf("%s/builds", c.BaseURL), nil)

	if req != nil {
		//built valid request
		t.Log("TEST PASSED: Request built successfully")
	}
}

func TestCreateRequestValidBody(t *testing.T) {
	c := NewClient("test", "test", "")

	//create a post body
	body := strings.NewReader("sample")

	//create a request with a body
	req := c.buildRequest("POST", fmt.Sprintf("%s/builds", c.BaseURL), body)

	if req != nil {
		t.Log("TEST PASSED: Request built successfully with a body")
	}
}

func TestClientBadRequest(t *testing.T) {
	c := NewClient("test", "test", "")
	err := c.sendRequest(http.NewRequest("POST", "", nil))
	if err != nil {
		t.Log("Sent bad request")
	}
}
