package goHttpProMaxPlus

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HttpMethod string

const (
	GET     HttpMethod = http.MethodGet
	POST    HttpMethod = http.MethodPost
	DELETE  HttpMethod = http.MethodDelete
	PUT     HttpMethod = http.MethodPut
	PATCH   HttpMethod = http.MethodPatch
	OPTIONS HttpMethod = http.MethodOptions
)

func (hm HttpMethod) String() string {
	switch hm {
	case GET, POST, DELETE, PUT, PATCH, OPTIONS:
		return string(hm)
	default:
		return "invalid http method"
	}
}

type HttpRequest struct {
	Method     HttpMethod
	URL        string
	Cookies    map[string]string
	Headers    map[string]string
	Params     map[string]string //追加URL中的RawQuery
	Forms      map[string]string
	Json       *string
	Xml        *string
	File       *os.File
	ReaderBody *io.Reader
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
	hr.Forms[name] = value
	return hr
}

func (hr *HttpRequest) AppendForms(forms map[string]string) *HttpRequest {
	for k, v := range forms {
		hr.AppendForm(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetForm(name string) string {
	return hr.Forms[name]
}

func (hr *HttpRequest) DelForm(name string) *HttpRequest {
	delete(hr.Forms, name)
	return hr
}

func (hr *HttpRequest) ClearAllForm() *HttpRequest {
	hr.Forms = map[string]string{}
	return hr
}

func (hr *HttpRequest) ReplaceForms(header map[string]string) *HttpRequest {
	hr.ClearAllForm()
	hr.AppendForms(header)
	return hr
}

func (hr *HttpRequest) SetJsonBody(json string) *HttpRequest {
	hr.Json = &json
	return hr
}

func (hr *HttpRequest) SetFileBody(file *os.File) *HttpRequest {
	hr.File = file
	return hr
}

func (hr *HttpRequest) SetFileBodyPath(path string) *HttpRequest {
	file, err := os.Open(path)
	if err != nil {
		hr.File = nil
	} else {
		hr.File = file
	}
	return hr
}

func (hr *HttpRequest) SetMethod(method HttpMethod) *HttpRequest {
	hr.Method = method
	return hr
}

func (hr *HttpRequest) SetURL(URL string) *HttpRequest {
	hr.URL = URL
	return hr
}

func (hr *HttpRequest) SetReaderBody(reader *io.Reader) *HttpRequest {
	hr.ReaderBody = reader
	return hr
}

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{
		Method:  GET,
		Cookies: map[string]string{},
		Headers: map[string]string{},
		Params:  map[string]string{},
		Forms:   map[string]string{},
	}
}

func (hr *HttpRequest) BuildRequest() (*http.Request, error) {
	parseUrl, err := url.Parse(hr.URL)
	if err != nil {
		return nil, errors.New("url parse err, msg: " + err.Error())
	}

	// append params
	if hr.Params != nil {
		parseUrl.RawQuery = parseParams(hr.Params)
	}

	// create body
	var body io.Reader
	if hr.Forms != nil {
		body = strings.NewReader(parseParams(hr.Forms))
	} else if hr.Json != nil {
		body = strings.NewReader(*hr.Json)
	} else if hr.Xml != nil {
		body = strings.NewReader(*hr.Xml)
	} else if hr.File != nil {
		body = hr.File
	} else if hr.ReaderBody != nil {
		body = *hr.ReaderBody
	} else {
		body = nil
	}

	request, err := http.NewRequest(hr.Method.String(), parseUrl.String(), body)
	if err != nil {
		return nil, errors.New("build request err, msg: " + err.Error())
	}

	// append header and cookie
	parseHeader(hr.Headers, request)
	parseCookie(hr.Cookies, request)

	return request, nil
}

func parseCookie(cookies map[string]string, request *http.Request) {
	for k, v := range cookies {
		request.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

func parseHeader(header map[string]string, request *http.Request) {
	for k, v := range header {
		request.Header.Set(k, v)
	}
}

func parseParams(param map[string]string) string {
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
