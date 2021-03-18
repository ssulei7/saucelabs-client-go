package sauce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func MockPositiveJobResponse(w http.ResponseWriter, r *http.Request) {
	resp := Jobs{
		Job{
			Status: "Completed",
		},
	}

	json.NewEncoder(w).Encode(&resp)
}

func MockPositiveJobResponseQueryParams(w http.ResponseWriter, r *http.Request) {
	resp := Jobs{
		Job{
			Status: "Completed",
		},
		Job{
			Status: "Failed",
		},
	}

	json.NewEncoder(w).Encode(&resp)
}

func MockNegativeJobResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}

func TestGetJob(t *testing.T) {

	//server a fake job back
	s := httptest.NewServer(http.HandlerFunc(MockPositiveJobResponse))

	//ensure we close connection
	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)
	val, err := c.GetJobs(nil)
	if err == nil {
		fmt.Println(val)
		t.Log("Test passed, got jobs for valid user")
	}
}

func TestGetJobQueryParams(t *testing.T) {
	//server a fake job back
	s := httptest.NewServer(http.HandlerFunc(MockPositiveJobResponseQueryParams))

	//ensure we close connection
	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)
	val, err := c.GetJobs(&JobRequestOptions{
		Limit: 2,
		Skip:  0,
	})
	if err == nil {
		t.Log("Test passed, got jobs for valid user")
	}

	if len(val) == 2 {
		t.Log("Correct amount of objects returned")
	}
}

func TestGetJobFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(MockNegativeJobResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)
	val, err := c.GetJobs(nil)
	if val == nil && err == nil {
		t.Log("Test passed, got an error and negative value")
	}
}
