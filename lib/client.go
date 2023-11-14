package doggl

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	baseURI    = "https://api.track.toggl.com/api/v9"
	retryCount = 3
	apiTokenKey ContextKey = "api_token"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	ctx    context.Context
	client HTTPClient
}

func NewClient(ctx context.Context, client HTTPClient) *Client {
	return &Client{
		ctx:    ctx,
		client: client,
	}
}

func NewDefaultClient(ctx context.Context) *Client {
	return &Client{
		ctx:    ctx,
		client: http.DefaultClient,
	}
}


func (client *Client) doRequest(method string, endpoint string, param interface{}) (res *http.Response, err error) {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, endpoint)

	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return nil, err
	}


	apiTokenValue := client.ctx.Value(apiTokenKey).(string)
	basic := base64.StdEncoding.EncodeToString([]byte(apiTokenValue + ":api_token"))
	req.Header.Add("Authorization", "Basic "+basic)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if param != nil {
		body, marshalErr := json.Marshal(param)
		if marshalErr != nil {
			return nil, marshalErr
		}

		req.Body = io.NopCloser(bytes.NewReader(body))
	}

	count := 0
	for count < retryCount {
		res, err := client.client.Do(req)
		if err == nil {
			return res, nil
		}
		count++
	}

	return nil, errors.New("Retry count exceeded.")
}
