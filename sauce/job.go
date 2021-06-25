package sauce

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Job struct {
	Status     string `json:"status"`
	BaseConfig struct {
		GoogChromeoptions struct {
			W3C        bool          `json:"w3c"`
			Args       []interface{} `json:"args"`
			Extensions []interface{} `json:"extensions"`
		} `json:"goog:chromeOptions"`
		SauceOptions   JobSauceOptions `json:"sauce:options"`
		Browsername    string          `json:"browserName"`
		Platformname   string          `json:"platformName"`
		Browserversion string          `json:"browserVersion"`
	} `json:"base_config"`
	CommandCounts      JobCommandCounts `json:"command_counts"`
	DeletionTime       interface{}      `json:"deletion_time"`
	URL                interface{}      `json:"url"`
	OrgID              string           `json:"org_id"`
	CreationTime       int              `json:"creation_time"`
	ID                 string           `json:"id"`
	TeamID             string           `json:"team_id"`
	PerformanceEnabled interface{}      `json:"performance_enabled"`
	AssignedTunnelID   interface{}      `json:"assigned_tunnel_id"`
	Container          bool             `json:"container"`
	GroupID            string           `json:"group_id"`
	Public             string           `json:"public"`
	Breakpointed       interface{}      `json:"breakpointed"`
}

type JobDetails struct {
	BrowserShortVersion   string        `json:"browser_short_version,omitempty"`
	VideoURL              string        `json:"video_url,omitempty"`
	CreationTime          int           `json:"creation_time,omitempty"`
	CustomData            interface{}   `json:"custom-data,omitempty"`
	BrowserVersion        string        `json:"browser_version,omitempty"`
	Owner                 string        `json:"owner,omitempty"`
	AutomationBackend     string        `json:"automation_backend,omitempty"`
	ID                    string        `json:"id,omitempty"`
	CollectsAutomatorLog  bool          `json:"collects_automator_log,omitempty"`
	RecordScreenshots     bool          `json:"record_screenshots,omitempty"`
	RecordVideo           bool          `json:"record_video,omitempty"`
	Build                 interface{}   `json:"build,omitempty"`
	Passed                interface{}   `json:"passed,omitempty"`
	Public                string        `json:"public,omitempty"`
	AssignedTunnelID      interface{}   `json:"assigned_tunnel_id,omitempty"`
	Status                string        `json:"status,omitempty"`
	LogURL                string        `json:"log_url,omitempty"`
	StartTime             int           `json:"start_time,omitempty"`
	Proxied               bool          `json:"proxied,omitempty"`
	ModificationTime      int           `json:"modification_time,omitempty"`
	Tags                  []interface{} `json:"tags,omitempty"`
	Name                  interface{}   `json:"name,omitempty"`
	CommandsNotSuccessful int           `json:"commands_not_successful,omitempty"`
	ConsolidatedStatus    string        `json:"consolidated_status,omitempty"`
	SeleniumVersion       interface{}   `json:"selenium_version,omitempty"`
	Manual                bool          `json:"manual,omitempty"`
	EndTime               int           `json:"end_time,omitempty"`
	Error                 string        `json:"error,omitempty"`
	Os                    string        `json:"os,omitempty"`
	Breakpointed          interface{}   `json:"breakpointed,omitempty"`
	Browser               string        `json:"browser,omitempty"`
}

type JobSauceOptions struct {
	Build string `json:"build"`
	Name  string `json:"name"`
}

type JobCommandCounts struct {
	All   int `json:"All"`
	Error int `json:"Error"`
}

type JobRequestOptions struct {
	Limit int
	Skip  int
}

type Jobs []Job

func (c *Client) GetJobs(options *JobRequestOptions) (Jobs, error) {
	limit := 100
	skip := 0

	if options != nil {
		limit = options.Limit
		skip = options.Skip
	}

	req := c.buildRequest("GET", fmt.Sprintf("%s/jobs", c.BaseURL), nil)

	//add options
	queryParams := req.URL.Query()
	queryParams.Add("limit", fmt.Sprint(limit))
	queryParams.Add("skip", fmt.Sprint(skip))
	req.URL.RawQuery = queryParams.Encode()

	res := Jobs{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetJobDetails(id string) (*JobDetails, error) {

	//build request
	req := c.buildRequest("GET", fmt.Sprintf("%s/jobs/%s", c.BaseURL, id), nil)

	res := &JobDetails{}

	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UpdateJob(id string, updatedDetails *JobDetails) (*JobDetails, error) {
	json, _ := json.Marshal(updatedDetails)

	req := c.buildRequest("PUT", fmt.Sprintf("%s/jobs/%s", c.BaseURL, id), bytes.NewReader(json))
	res := &JobDetails{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}
	return res, nil
}
