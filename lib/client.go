package doggl

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	baseURI    = "https://api.track.toggl.com/api/v9"
	retryCount = 3
	apiToken ContextKey = "api_token"
)


type Client struct {
	ctx    context.Context
	client *http.Client
}

func NewClient(ctx context.Context, client *http.Client) *Client {
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

func (client *Client) do(method string, endpoint string, param interface{}) (res *http.Response, err error) {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, endpoint)

	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	basic := base64.StdEncoding.EncodeToString([]byte(client.ctx.Value(apiToken).(string) + ":api_token"))
	req.Header.Add("Authorization", "Basic "+basic)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if param != nil {
		body, marshalErr := json.Marshal(param)
		if marshalErr != nil {
			return nil, marshalErr
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body))
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
