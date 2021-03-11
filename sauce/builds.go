package sauce

import (
	"fmt"
	"net/http"
)

//Build .. information of a sauce build
type Build struct {
	Status           string      `json:"status"`
	Jobs             BuildJob    `json:"jobs"`
	Name             string      `json:"name"`
	DeletionTime     interface{} `json:"deletion_time"`
	OrgID            string      `json:"org_id"`
	StartTime        int         `json:"start_time"`
	CreationTime     int         `json:"creation_time"`
	Number           interface{} `json:"number"`
	Public           bool        `json:"public"`
	ModificationTime int         `json:"modification_time"`
	Prefix           interface{} `json:"prefix"`
	EndTime          int         `json:"end_time"`
	Passed           interface{} `json:"passed"`
	Owner            string      `json:"owner"`
	Run              int         `json:"run"`
	TeamID           string      `json:"team_id"`
	GroupID          string      `json:"group_id"`
	ID               string      `json:"id"`
}

//BuildJob ...
type BuildJob struct {
	Completed int `json:"completed"`
	Finished  int `json:"finished"`
	Queued    int `json:"queued"`
	Failed    int `json:"failed"`
	Running   int `json:"running"`
	Passed    int `json:"passed"`
	Errored   int `json:"errored"`
	Public    int `json:"public"`
}

//Builds list of sauce builds
type Builds []Build

//GetBuilds get all builds for specified user
func (c *Client) GetBuilds() (*Builds, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/builds", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	resp := Builds{}

	if err := c.sendRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
