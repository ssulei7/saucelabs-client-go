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

func MockNegativeJobResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}

func TestGetJob(t *testing.T) {

	//server a fake job back
	s := httptest.NewServer(http.HandlerFunc(MockPositiveJobResponse))

	//ensure we close connection
	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "ssuleiman", s.URL)
	val, err := c.GetJobs()
	if err == nil {
		fmt.Println(val)
		t.Log("Test passed, got jobs for valid user")
	}
}

func TestGetJobFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(MockNegativeJobResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "ssuleiman", s.URL)
	val, err := c.GetJobs()
	if val == nil && err == nil {
		t.Log("Test passed, got an error and negative value")
	}
}
