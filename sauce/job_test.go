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

func MockPositiveJobDetailsResponse(w http.ResponseWriter, r *http.Request) {
	resp := JobDetails{
		BrowserShortVersion:   "90",
		VideoURL:              "someurl",
		CreationTime:          1619124503,
		CustomData:            nil,
		BrowserVersion:        "90.0.4430.72",
		Owner:                 "testuser",
		AutomationBackend:     "webdriver",
		ID:                    "b3d75978ab5949c9b5adb7dad4317fd0",
		CollectsAutomatorLog:  false,
		RecordScreenshots:     true,
		RecordVideo:           true,
		Build:                 nil,
		Passed:                nil,
		Public:                "team",
		AssignedTunnelID:      nil,
		Status:                "complete",
		LogURL:                "someurl",
		StartTime:             1619124503,
		Proxied:               false,
		ModificationTime:      1619124642,
		Tags:                  nil,
		Name:                  nil,
		CommandsNotSuccessful: 1,
		ConsolidatedStatus:    "error",
		SeleniumVersion:       nil,
		Manual:                false,
		EndTime:               1619124642,
		Error:                 "Test did not see a new command for 90 seconds. Timing out.",
		Os:                    "Windows 10",
		Breakpointed:          nil,
		Browser:               "googlechrome",
	}

	json.NewEncoder(w).Encode(&resp)
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

func TestGetJobDetails(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(MockPositiveJobDetailsResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)

	val, err := c.GetJobDetails("123456")

	if err == nil && val != nil {
		t.Log("Got a resp back from the endpoint")
	}

	//validate some data
	if val.Browser == "googlechrome" {
		t.Log("Got some job data")
	}
}

func TestGetJobDetailsFail(t *testing.T) {
	//have some invalid response that doesn't conform to contract
	s := httptest.NewServer(http.HandlerFunc(MockPositiveJobResponse))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)

	val, err := c.GetJobDetails("123456")

	if err != nil && val == nil {
		t.Log("Got a resp back from the endpoint that was invalid")
	}
}

func TestUpdateJob(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		update := JobDetails{}
		json.NewDecoder(r.Body).Decode(&update)
		resp := JobDetails{
			BrowserShortVersion:   "90",
			VideoURL:              update.VideoURL,
			CreationTime:          1619124503,
			CustomData:            nil,
			BrowserVersion:        "90.0.4430.72",
			Owner:                 "testuser",
			AutomationBackend:     "webdriver",
			ID:                    "b3d75978ab5949c9b5adb7dad4317fd0",
			CollectsAutomatorLog:  false,
			RecordScreenshots:     true,
			RecordVideo:           true,
			Build:                 nil,
			Passed:                nil,
			Public:                "team",
			AssignedTunnelID:      nil,
			Status:                "complete",
			LogURL:                "someurl",
			StartTime:             1619124503,
			Proxied:               false,
			ModificationTime:      1619124642,
			Tags:                  nil,
			Name:                  nil,
			CommandsNotSuccessful: 1,
			ConsolidatedStatus:    "error",
			SeleniumVersion:       nil,
			Manual:                false,
			EndTime:               1619124642,
			Error:                 "Test did not see a new command for 90 seconds. Timing out.",
			Os:                    "Windows 10",
			Breakpointed:          nil,
			Browser:               "googlechrome",
		}

		json.NewEncoder(rw).Encode(&resp)
	}))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)

	updatedJobDetails := JobDetails{
		VideoURL: "stuff",
	}

	val, err := c.UpdateJob("123", &updatedJobDetails)

	if err == nil && val != nil {
		t.Log("Approrpriate values returned")
	}

}

func TestUpdateJobFail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Not able to update job"))
		rw.WriteHeader(http.StatusBadRequest)
	}))

	defer s.Close()

	c := NewClient(os.Getenv("SAUCE_KEY"), "test", s.URL)

	updatedJobDetails := JobDetails{
		VideoURL: "someurl",
	}

	val, err := c.UpdateJob("123", &updatedJobDetails)

	if err != nil && val == nil {
		t.Log("Successfully failed")
	}

}
