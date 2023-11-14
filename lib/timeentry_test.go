package doggl

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

type MockHTTPClientWithResponse struct{}

func (m *MockHTTPClientWithResponse) Do(req *http.Request) (*http.Response, error) {
	TimeEntryResponse := TimeEntryResponse{
		ID:              0,
		ProjectID:       123456789,
		Duration:        100,
		Description:     "Test Description",
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
		Description: "Test Description",
		Duration:    100,
		ProjectId:   123456789,
	}

	res, err := client.StartTimeEntry(timeEntry)
	assert.Equal(t, nil, err)
	assert.Equal(t, 100, res.Duration)
}
