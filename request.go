package goHttpProMaxPlus

import (
	"bytes"
	"errors"
	sj "github.com/bitly/go-simplejson"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

func (hm HttpMethod) String() string {
	switch hm {
	case GET:
		return "GET"
	case POST:
		return "POST"
	default:
		return "GET"
	}
}

type HttpRequest struct {
	Method    HttpMethod
	URL       *url.URL
	Cookies   map[string]string
	Headers   map[string]string
	Params    map[string]string //覆盖URL中的RawQuery
	FormsBody map[string]string
	JsonBody  *sj.Json
	FileBody  *os.File
}

func (hr *HttpRequest) AppendCookie(name, value string) *HttpRequest {
	hr.Cookies[name] = value
	return hr
}

func (hr *HttpRequest) AppendCookies(cookies map[string]string) *HttpRequest {
	for k, v := range cookies {
		hr.AppendCookie(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetCookie(name string) string {
	return hr.Cookies[name]
}

func (hr *HttpRequest) DelCookie(name string) *HttpRequest {
	delete(hr.Cookies, name)
	return hr
}

func (hr *HttpRequest) ClearAllCookie() *HttpRequest {
	hr.Cookies = map[string]string{}
	return hr
}

func (hr *HttpRequest) ReplaceCookies(cookies map[string]string) {
	hr.ClearAllCookie()
	for k, v := range cookies {
		hr.AppendCookie(k, v)
	}
}

func (hr *HttpRequest) AppendParam(name, value string) *HttpRequest {
	hr.Params[name] = value
	return hr
}

func (hr *HttpRequest) AppendParams(params map[string]string) *HttpRequest {
	for k, v := range params {
		hr.AppendParam(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetParam(name string) string {
	return hr.Params[name]
}

func (hr *HttpRequest) DelParam(name string) *HttpRequest {
	delete(hr.Params, name)
	return hr
}

func (hr *HttpRequest) ClearAllParam() *HttpRequest {
	hr.Params = map[string]string{}
	return hr
}

func (hr *HttpRequest) ReplaceParams(param map[string]string) *HttpRequest {
	hr.ClearAllParam()
	hr.AppendParams(param)
	return hr
}

func (hr *HttpRequest) AppendHeader(name, value string) *HttpRequest {
	hr.Headers[name] = value
	return hr
}

func (hr *HttpRequest) AppendHeaders(headers map[string]string) *HttpRequest {
	for k, v := range headers {
		hr.AppendHeader(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetHeader(name string) string {
	return hr.Headers[name]
}

func (hr *HttpRequest) DelHeader(name string) *HttpRequest {
	delete(hr.Headers, name)
	return hr
}

func (hr *HttpRequest) ClearAllHeader() *HttpRequest {
	hr.Headers = map[string]string{}
	return hr
}

func (hr *HttpRequest) ReplaceHeaders(header map[string]string) *HttpRequest {
	hr.ClearAllHeader()
	hr.AppendHeaders(header)
	return hr
}

func (hr *HttpRequest) AppendForm(name, value string) *HttpRequest {
	hr.FormsBody[name] = value
	return hr
}

func (hr *HttpRequest) AppendForms(forms map[string]string) *HttpRequest {
	for k, v := range forms {
		hr.AppendForm(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetForm(name string) string {
	return hr.FormsBody[name]
}

func (hr *HttpRequest) DelForm(name string) *HttpRequest {
	delete(hr.FormsBody, name)
	return hr
}

func (hr *HttpRequest) ClearAllForm() *HttpRequest {
	hr.FormsBody = map[string]string{}
	return hr
}

func (hr *HttpRequest) ReplaceForms(header map[string]string) *HttpRequest {
	hr.ClearAllForm()
	hr.AppendForms(header)
	return hr
}

func (hr *HttpRequest) SetJsonBodySJ(json *sj.Json) *HttpRequest {
	hr.JsonBody = json
	return hr
}

func (hr *HttpRequest) SetJsonBodyStr(json string) *HttpRequest {
	newJson, err := sj.NewJson([]byte(json))
	if err != nil {
		hr.JsonBody = sj.New()
	} else {
		hr.JsonBody = newJson
	}
	return hr
}

func (hr *HttpRequest) SetFileBody(file *os.File) *HttpRequest {
	hr.FileBody = file
	return hr
}

func (hr *HttpRequest) SetFileBodyPath(path string) *HttpRequest {
	file, err := os.Open(path)
	if err != nil {
		hr.FileBody = nil
	} else {
		hr.FileBody = file
	}
	return hr
}

func GetHttpRequest(method HttpMethod, url *url.URL) *HttpRequest {
	return &HttpRequest{
		Method:    method,
		URL:       url,
		Cookies:   map[string]string{},
		Headers:   map[string]string{},
		Params:    map[string]string{},
		FormsBody: map[string]string{},
	}
}

func BuildRequest(hr *HttpRequest) (*http.Request, error) {
	var body io.Reader

	if hr.Params != nil {
		hr.URL.RawQuery = buildParams(hr.Params)
	}

	if hr.FormsBody != nil {
		body = strings.NewReader(buildParams(hr.Params))
	} else if hr.JsonBody != nil {
		json, err := hr.JsonBody.MarshalJSON()
		if err != nil {
			return nil, errors.New("json err, msg: " + err.Error())
		}
		body = bytes.NewReader(json)
	} else if hr.FileBody != nil {
		body = hr.FileBody
	} else {
		body = nil
	}

	request, err := http.NewRequest(hr.Method.String(), hr.URL.String(), body)
	if err != nil {
		return nil, errors.New("build request err, msg: " + err.Error())
	}

	return request, nil
}

func buildParams(param map[string]string) string {
	var buff strings.Builder

	for k, v := range param {
		if buff.Len() > 0 {
			buff.WriteByte('&')
		}
		buff.WriteString(url.QueryEscape(k))
		buff.WriteByte('=')
		buff.WriteString(url.QueryEscape(v))
	}

	return buff.String()
}
