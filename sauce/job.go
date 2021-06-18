package sauce

import (
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
	BrowserShortVersion   string        `json:"browser_short_version"`
	VideoURL              string        `json:"video_url"`
	CreationTime          int           `json:"creation_time"`
	CustomData            interface{}   `json:"custom-data"`
	BrowserVersion        string        `json:"browser_version"`
	Owner                 string        `json:"owner"`
	AutomationBackend     string        `json:"automation_backend"`
	ID                    string        `json:"id"`
	CollectsAutomatorLog  bool          `json:"collects_automator_log"`
	RecordScreenshots     bool          `json:"record_screenshots"`
	RecordVideo           bool          `json:"record_video"`
	Build                 interface{}   `json:"build"`
	Passed                interface{}   `json:"passed"`
	Public                string        `json:"public"`
	AssignedTunnelID      interface{}   `json:"assigned_tunnel_id"`
	Status                string        `json:"status"`
	LogURL                string        `json:"log_url"`
	StartTime             int           `json:"start_time"`
	Proxied               bool          `json:"proxied"`
	ModificationTime      int           `json:"modification_time"`
	Tags                  []interface{} `json:"tags"`
	Name                  interface{}   `json:"name"`
	CommandsNotSuccessful int           `json:"commands_not_successful"`
	ConsolidatedStatus    string        `json:"consolidated_status"`
	SeleniumVersion       interface{}   `json:"selenium_version"`
	Manual                bool          `json:"manual"`
	EndTime               int           `json:"end_time"`
	Error                 string        `json:"error"`
	Os                    string        `json:"os"`
	Breakpointed          interface{}   `json:"breakpointed"`
	Browser               string        `json:"browser"`
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
