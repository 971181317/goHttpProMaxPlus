package goHttpProMaxPlus

import (
	"errors"
	"net/http"
	"time"
)

// HttpClient 请求http的客户端
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

	// todo 处理response

	client.AfterResponseCreate(client.AspectArgs)

	return &HttpResponse{
		res: _resp,
	}, nil
}

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
