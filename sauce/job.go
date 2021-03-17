package sauce

import (
	"fmt"
	"net/http"
)

type Job struct {
	Status     string `json:"status"`
	BaseConfig struct {
		GoogChromeoptions struct {
			W3C        bool          `json:"w3c"`
			Args       []interface{} `json:"args"`
			Extensions []interface{} `json:"extensions"`
		} `json:"goog:chromeOptions"`
		SauceOptions struct {
			Build string `json:"build"`
			Name  string `json:"name"`
		} `json:"sauce:options"`
		Browsername    string `json:"browserName"`
		Platformname   string `json:"platformName"`
		Browserversion string `json:"browserVersion"`
	} `json:"base_config"`
	CommandCounts struct {
		All   int `json:"All"`
		Error int `json:"Error"`
	} `json:"command_counts"`
	DeletionTime       interface{} `json:"deletion_time"`
	URL                interface{} `json:"url"`
	OrgID              string      `json:"org_id"`
	CreationTime       int         `json:"creation_time"`
	ID                 string      `json:"id"`
	TeamID             string      `json:"team_id"`
	PerformanceEnabled interface{} `json:"performance_enabled"`
	AssignedTunnelID   interface{} `json:"assigned_tunnel_id"`
	Container          bool        `json:"container"`
	GroupID            string      `json:"group_id"`
	Public             string      `json:"public"`
	Breakpointed       interface{} `json:"breakpointed"`
}

type Jobs []Job

func (c *Client) GetJobs() (Jobs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/jobs", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := Jobs{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
