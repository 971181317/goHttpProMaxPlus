package goHttpProMaxPlus

import (
	"errors"
	"io"
	"net/http"
	"time"
)

// AspectModel 切片模组
type AspectModel func(...interface{})

// HttpClient A client send request
type HttpClient struct {
	c                   *http.Client
	BeforeClientBuild   AspectModel
	AfterClientBuild    AspectModel
	BeforeRequestBuild  AspectModel
	AfterRequestBuild   AspectModel
	AfterResponseCreate AspectModel
	AspectArgs          []interface{} // Only the most recent assignment will be kept
}

func (client HttpClient) GetAspectArgs() []interface{} {
	return client.AspectArgs
}

func (client HttpClient) Get(url string) (*HttpResponse, error) {
	return client.GetWithCookieAndHeader(url, nil, nil)
}

func (client HttpClient) GetWithCookieAndHeader(url string, cookie, header map[string]string) (*HttpResponse, error) {
	return client.Do(&HttpRequest{
		Method: GET,
		URL: url,
		Headers: header,
		Cookies: cookie,
	})
}

func (client HttpClient) Post(url string) (*HttpResponse, error) {
	return client.PostWithCookieHeaderAndForm(url, nil, nil, nil)
}

func (client HttpClient) PostWithForm(url string, form map[string]string) (*HttpResponse, error) {
	return client.PostWithCookieHeaderAndForm(url, nil, nil, form)
}

func (client HttpClient) PostWithCookieHeaderAndForm(url string, cookie, header, form map[string]string) (*HttpResponse, error) {
	return client.Do(&HttpRequest{
		Method: POST,
		URL: url,
		Headers: header,
		Cookies: cookie,
		Forms: form,
	})
}

func (client HttpClient) PostWithIoData(url string, data *io.Reader) (*HttpResponse, error) {
	return client.PostWithCookieHeaderAndIoData(url, nil, nil, data)
}

func (client HttpClient) PostWithCookieHeaderAndIoData(url string, cookie, header map[string]string, data *io.Reader) (*HttpResponse, error) {
	return client.Do(&HttpRequest{
		URL: url,
		Headers: header,
		Cookies: cookie,
		ReaderBody: data,
	})
}

// Do run with Aspect
func (client HttpClient) Do(req *HttpRequest) (*HttpResponse, error) {
	client.BeforeRequestBuild(client.AspectArgs)

	_req, err := req.BuildRequest()
	if err != nil {
		return nil, err
	}

	client.AfterRequestBuild(client.AspectArgs)

	_resp, err := client.c.Do(_req)
	if err != nil {
		return nil, errors.New("request err, msg : " + err.Error())
	}

	resp := CreateResponse(_resp)

	client.AfterResponseCreate(client.AspectArgs)

	return resp, nil
}

// defaultAspect Default aspect
func defaultAspect(...interface{}) {}

var DefaultClient = &HttpClient{
	&http.Client{Timeout: 5 * time.Second},
	defaultAspect,
	defaultAspect,
	defaultAspect,
	defaultAspect,
	defaultAspect,
	nil,
}

// GetDefaultClient
// !!! This method will not execute BeforeClientBuild and AfterClientBuild.
// 这个方法不会执行 BeforeClientBuild 和 AfterClientBuild。
func GetDefaultClient() *HttpClient {
	return DefaultClient
}

// NewClient
// !!! This method will not execute BeforeClientBuild and AfterClientBuild.
// 这个方法不会执行 BeforeClientBuild 和 AfterClientBuild。
func NewClient(client *http.Client) *HttpClient {
	return NewClientX(client, nil, nil, nil, nil, nil)
}

// NewClientX
// This method execute BeforeClientBuild and AfterClientBuild.
// 这个方法执行 BeforeClientBuild 和 AfterClientBuild。
func NewClientX(client *http.Client,
	beforeClientBuild, afterClientBuild, beforeRequestBuild, afterRequestBuild, afterResponseCreate AspectModel,
	args ...interface{}) *HttpClient {
	beforeRequestBuild(args)

	hc := &HttpClient{
		c:                   client,
		BeforeClientBuild:   beforeClientBuild,
		AfterClientBuild:    afterClientBuild,
		BeforeRequestBuild:  beforeRequestBuild,
		AfterRequestBuild:   afterRequestBuild,
		AfterResponseCreate: afterResponseCreate,
		AspectArgs:          args,
	}

	afterRequestBuild(args)

	return hc
}
