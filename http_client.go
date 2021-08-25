package goHttpProMaxPlus

import "net/http"

type httpClient struct {
	client *http.Client
}

func GetDefaultClient() *httpClient {
	return &httpClient{
		client: &http.Client{},
	}
}
