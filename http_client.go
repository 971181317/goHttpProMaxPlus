package goHttpProMaxPlus

import (
	sj "github.com/bitly/go-simplejson"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AspectModel 切片模组
type AspectModel func(...interface{})

// HttpClient 请求http的客户端
type HttpClient struct {
	client *http.Client
	Before AspectModel //请求之前的操作
	After  AspectModel //请求后的操作
	Error  AspectModel //出现错误的操作
}

func (hc *HttpClient) NewStringBodyRequest(method, url string, body string) (*HttpRequest, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	return &HttpRequest{
		request: req,
	}, nil
}

func (hc *HttpClient) NewJsonBodyRequest(method, url string, body *sj.Json) (*HttpRequest, error) {
	json, err := body.String()
	if err != nil {
		return nil, err
	}
	return hc.NewStringBodyRequest(method, url, json)
}

func (hc *HttpClient) NewGetJsonBodyRequest(url string, body *sj.Json) (*HttpRequest, error) {
	return hc.NewJsonBodyRequest("Get", url, body)
}

func (hc *HttpClient) NewPostJsonBodyRequest(url string, body *sj.Json) (*HttpRequest, error) {
	return hc.NewJsonBodyRequest("Post", url, body)
}

func (hc *HttpClient) NewFormBodyRequest(method, url string, form url.Values) (*HttpRequest, error) {
	return hc.NewStringBodyRequest(method, url, form.Encode())
}

func (hc *HttpClient) NewGetFormBodyRequest(url string, form url.Values) (*HttpRequest, error) {
	return hc.NewFormBodyRequest("Get", url, form)
}

func (hc *HttpClient) NewPostFormBodyRequest(url string, form url.Values) (*HttpRequest, error) {
	return hc.NewFormBodyRequest("Post", url, form)
}

func (hc *HttpClient) NewRequest(method, url string) (*HttpRequest, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return &HttpRequest{
		request: req,
	}, nil
}

func (hc *HttpClient) NewGetRequest(method, url string, body io.Reader) (*HttpRequest, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	return &HttpRequest{
		request: req,
	}, nil
}

func NewUserClient(client *http.Client) *HttpClient {
	return NewUserClientWithAspect(client, nil, nil, nil)
}

func NewClient() *HttpClient {
	return NewUserClientWithAspect(http.DefaultClient, nil, nil, nil)
}

func NewUserClientWithAspect(client *http.Client, before, after, error AspectModel) *HttpClient {
	return &HttpClient{
		client: client,
		Before: before,
		After:  after,
		Error:  error,
	}
}
