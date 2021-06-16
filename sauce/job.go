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
