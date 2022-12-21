package doggl

import (
	"bytes"
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
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(client *http.Client, token string) *Client {
	return &Client{
		client: client,
		token:  token,
	}
}

func NewDefaultClient(token string) *Client {
	return &Client{
		client: http.DefaultClient,
		token:  token,
	}
}

func (client *Client) do(method string, endpoint string, param interface{}) (res *http.Response, err error) {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, endpoint)

	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return
	}

	basic := base64.StdEncoding.EncodeToString([]byte(client.token + ":api_token"))
	req.Header.Add("Authorization", "Basic "+basic)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if param != nil {
		body, marshalErr := json.Marshal(param)
		if marshalErr != nil {
			return
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

	count := 0
	for count < retryCount {
		res, err := client.client.Do(req)
		if err == nil {
			return res, err
		}
		count++
	}

	return nil, errors.New("Retry count exceeded.")
}
