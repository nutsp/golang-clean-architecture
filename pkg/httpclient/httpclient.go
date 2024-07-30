package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nutsp/golang-clean-architecture/config"
)

var (
	MethodGet    string = http.MethodGet
	MethodPost   string = http.MethodPost
	MethodPut    string = http.MethodPut
	MethodDelete string = http.MethodDelete
	MethodPatch  string = http.MethodPatch
)

var (
	ContentType     string = "ContentType"
	ApplicationJson string = "application/json"
)

type IClient interface {
	Do(ctx context.Context, req *Request) (*Response, error)
}

type Header map[string]string

type Request struct {
	Method    string
	URL       string
	Header    Header
	Body      interface{}
	bodyBytes []byte
}

type Response struct {
	*http.Response
	Body []byte
}

type Client struct {
	client *http.Client
}

func NewClient(cfg config.HttpClient) *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (req *Request) newRequest() (*http.Request, error) {
	if req.Body != nil {
		req.bodyBytes, _ = json.Marshal(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, req.URL, bytes.NewReader(req.bodyBytes))
	if err != nil {
		return nil, err
	}

	for k, v := range req.Header {
		httpReq.Header.Set(k, v)
	}

	return httpReq, nil
}

func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	httpReq, err := req.newRequest()
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Response: res,
		Body:     body,
	}, nil
}

func (res *Response) IsSuccess() bool {
	return res.StatusCode >= 200 && res.StatusCode < 300
}
