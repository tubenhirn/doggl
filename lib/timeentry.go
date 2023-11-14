package doggl

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TimeEntry struct {
	Billable    bool   `json:"billable,omitempty"`
	CreatedWith string `json:"created_with,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    int64  `json:"duration,omitempty"`
	Duronly     bool   `json:"duronly,omitempty"`
	ProjectId   int    `json:"project_id,omitempty"`
	Start       string `json:"start,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	Stop        string `json:"stop,omitempty"`
	TaskId      int    `json:"task_id,omitempty"`
	WorkspaceId int    `json:"workspace_id,omitempty"`
}

func (client *Client) StartTimeEntry(timeEntry TimeEntry) (response TimeEntryResponse, err error) {
	endpoint := fmt.Sprintf("/workspaces/%d/time_entries", timeEntry.WorkspaceId)
	res, err := client.doRequest("POST", endpoint, timeEntry)
	if err != nil {
		return TimeEntryResponse{}, err
	}

	if res.StatusCode == 200 {
		enc := json.NewDecoder(res.Body)
		if err := enc.Decode(&response); err != nil {
			return TimeEntryResponse{}, err
		}
		return response, nil
	} else {
		return TimeEntryResponse{}, errors.New("request failed - " + res.Status)
	}
}
