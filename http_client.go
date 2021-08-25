package goHttpProMaxPlus

import (
	"net/http"
)

// HttpClient 请求http的客户端
type HttpClient struct {
	c                   *http.Client
	BeforeClientBuild   AspectModel
	AfterClientBuild    AspectModel
	BeforeRequestBuild  AspectModel
	AfterRequestBuild   AspectModel
	AfterResponseCreate AspectModel
}

func NewClientUserConf(client *http.Client) *HttpClient {
	return NewUserClientWithUserConfAndAspect(client, nil, nil, nil, nil, nil)
}

func NewClient() *HttpClient {
	return NewUserClientWithUserConfAndAspect(http.DefaultClient, nil, nil, nil, nil, nil)
}

func NewUserClientWithUserConfAndAspect(client *http.Client,
	beforeClientBuild, afterClientBuild, beforeRequestBuild, afterRequestBuild, afterResponseCreate AspectModel) *HttpClient {
	return &HttpClient{
		c:                   client,
		BeforeClientBuild:   beforeClientBuild,
		AfterClientBuild:    afterClientBuild,
		BeforeRequestBuild:  beforeRequestBuild,
		AfterRequestBuild:   afterRequestBuild,
		AfterResponseCreate: afterResponseCreate,
	}
}
