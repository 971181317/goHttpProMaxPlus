package goHttpProMaxPlus

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	resp    *http.Response
	Headers map[string]string
	Body    io.ReadCloser
	bodyStr *string
}

func (hr *HttpResponse) GetRespStr() string {
	if hr.bodyStr == nil {
		hr.setBodyStr()
	}

	return *hr.bodyStr
}

// GetResponse It return http.Response struct
func (hr *HttpResponse) GetResponse() *http.Response {
	return hr.resp
}

func (hr *HttpResponse) ParseJson(v interface{}) {
	_ = json.Unmarshal([]byte(hr.GetRespStr()), &v)
}

func CreateResponse(resp *http.Response) *HttpResponse{
	response := &HttpResponse{
		resp: resp,
	}

	response.parseHeader()
	response.parseBody()

	return response
}

func (hr *HttpResponse) setBodyStr() {
	defer hr.Body.Close()

	body, _ := ioutil.ReadAll(hr.Body)

	_s := string(body)
	hr.bodyStr = &_s
}

func (hr *HttpResponse) parseHeader() {
	hr.Headers = map[string]string{}

	header := hr.resp.Header
	for k, _ := range header {
		hr.Headers[k] = header.Get(k)
	}
}

func (hr *HttpResponse) parseBody() {
	hr.Body = hr.resp.Body
}
