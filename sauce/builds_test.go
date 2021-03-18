package sauce

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func MockPositiveBuildResponse(w http.ResponseWriter, r *http.Request) {
	resp := Builds{
		Build{
			Status: "complete",
			Jobs: BuildJob{
				Completed: 1,
				Finished:  1,
				Queued:    0,
				Failed:    0,
				Running:   0,
				Passed:    0,
				Errored:   0,
				Public:    0,
			},
			Name:             "ExampleBuild",
			DeletionTime:     nil,
			OrgID:            "",
			StartTime:        1614839114,
			CreationTime:     1614839137,
			Number:           nil,
			Public:           false,
			ModificationTime: 1614839150,
			Prefix:           nil,
			EndTime:          1614839125,
			Passed:           nil,
			Owner:            "testuser",
			Run:              0,
			TeamID:           "",
			GroupID:          "",
			ID:               "123456",
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func MockNegativeBuildResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func TestGetBuilds(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(MockPositiveBuildResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "user", s.URL)
	builds, err := c.GetBuilds()
	if err == nil {
		t.Log("Was able to get builds")
	}

	if builds[0].Name == "ExampleBuild" {
		t.Log("Got correct build name in slice")
	}
}

func TestGetBuildsWrongUser(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(MockNegativeJobResponse))

	defer s.Close()
	c := NewClient(os.Getenv("SAUCE_KEY"), "randomuser", s.URL)
	_, err := c.GetBuilds()
	if err != nil {
		t.Log("Test passed, wrong user")
	}
}

func TestGenerateBuildURL(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(MockPositiveBuildResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "user", s.URL)
	builds, err := c.GetBuilds()
	if err == nil {
		t.Log("Was able to get builds")
	}

	if len(builds) >= 1 {
		t.Log("Have more than one build!")
	}

	url := builds[0].GenerateBuildURL()
	if url == "app.saucelabs.com/builds/123456" {
		t.Log("URL Generated correctly")
	}
}

func TestGenerateBuildURLFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(MockPositiveBuildResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "user", s.URL)
	builds, err := c.GetBuilds()
	if err == nil {
		t.Log("Was able to get builds")
	}

	if len(builds) >= 1 {
		t.Log("Have more than one build!")
	}

	//change id to empty string
	builds[0].ID = ""

	url := builds[0].GenerateBuildURL()
	if url != "app.saucelabs.com/builds/123456" {
		t.Log("URL failed to generate")
	}
}
