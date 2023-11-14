package doggl

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

type MockHTTPClientWithResponse struct{}

func (m *MockHTTPClientWithResponse) Do(req *http.Request) (*http.Response, error) {
	TimeEntryResponse := TimeEntryResponse{
		ID:              0,
		WorkspaceID:     0,
		ProjectID:       0,
		TaskID:          nil,
		Billable:        false,
		Start:           time.Time{},
		Stop:            time.Time{},
		Duration:        100,
		Description:     "",
		Tags:            nil,
		TagIds:          nil,
		Duronly:         false,
		At:              time.Time{},
		ServerDeletedAt: nil,
		UserID:          0,
		UID:             0,
		Wid:             0,
		Pid:             0,
	}
	entry, _ := json.Marshal(&TimeEntryResponse)
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer(entry)),
		Status: "200 OK",
		StatusCode: 200,
	}

	return response, nil
}

func TestStartTimeEntry(t *testing.T){
	httpClient := &MockHTTPClientWithResponse{}
	ctx := context.WithValue(context.Background(), apiTokenKey, "test")
	client := NewClient(ctx, httpClient)

	timeEntry := TimeEntry{
		Billable:    false,
		CreatedWith: "",
		Description: "Test Description",
		Duration:    100,
		Duronly:     false,
		ProjectId:   123456789,
		Start:       "",
		StartDate:   "",
		Stop:        "",
		TaskId:      0,
		WorkspaceId: 0,
	}

	res, err := client.StartTimeEntry(timeEntry)
	assert.Equal(t, nil, err)
	assert.Equal(t, 100, res.Duration)
}
