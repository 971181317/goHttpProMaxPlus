package goHttpProMaxPlus

import (
	"encoding/json"
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

// HttpRequest
// Forms > Json > Xml > File > ReaderBody
type HttpRequest struct {
	Method     HttpMethod

	// Queries will overwrite the queries in the URL
	// Queries 会覆盖 URL 中的参数
	URL        string

	Cookies    map[string]string
	Headers    map[string]string

	// Queries will overwrite the queries in the URL
	// Queries 会覆盖 URL 中的参数
	Queries    map[string]string

	// When creating the HttpRequest, we have already added "Content-Type: application/x-www-form-urlencoded" for you .
	// 在创建 HttpRequest 的时候，我们已经帮你加好了 Content-Type: application/x-www-form-urlencoded
	// If you want to modify, please add "Content-Type" sauce to the Headers
	// 如果你想修改的话，请在 Headers 中加入“Content-Type”
	// The HTTP client ignores Form and uses Body instead.
	Forms      map[string]string

	// When creating the HttpRequest, we have already added "Content-Type: application/json" for you .
	// 在创建 HttpRequest 的时候，我们已经帮你加好了 Content-Type
	// If you want to modify, please add "Content-Type" sauce to the Headers
	// 如果你想修改的话，请在 Headers 中加入“Content-Type”
	Json       *string

	// When creating the HttpRequest, we have already added "Content-Type: text/xml" for you .
	// 在创建 HttpRequest 的时候，我们已经帮你加好了 Content-Type: application/x-www-form-urlencoded
	// If you want to modify, please add "Content-Type" sauce to the Headers
	// 如果你想修改的话，请在 Headers 中加入“Content-Type”
	Xml        *string

	// When creating the HttpRequest, we have already added "Content-Type: multipart/form-data" for you .
	// 在创建 HttpRequest 的时候，我们已经帮你加好了 Content-Type
	// If you want to modify, please add "Content-Type: " sauce to the Headers
	// 如果你想修改的话，请在 Headers 中加入“Content-Type”
	File       *os.File

	// Please add "Content-Type" sauce to the Headers
	// 请在 Headers 中加入“Content-Type”
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

func (hr *HttpRequest) AppendQuery(name, value string) *HttpRequest {
	hr.Queries[name] = value
	return hr
}

func (hr *HttpRequest) AppendQueries(params map[string]string) *HttpRequest {
	for k, v := range params {
		hr.AppendQuery(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetQuery(name string) string {
	return hr.Queries[name]
}

func (hr *HttpRequest) DelQuery(name string) *HttpRequest {
	delete(hr.Queries, name)
	return hr
}

func (hr *HttpRequest) ClearAllQuery() *HttpRequest {
	hr.Queries = map[string]string{}
	return hr
}

func (hr *HttpRequest) ReplaceQuery(param map[string]string) *HttpRequest {
	hr.ClearAllQuery()
	hr.AppendQueries(param)
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

func (hr *HttpRequest) SetJsonBodyStruct(v interface{}) *HttpRequest {
	marshal, err := json.Marshal(v)
	if err != nil {
		return hr
	}
	hr.Json = strPrt(string(marshal))
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
		Queries: map[string]string{},
		Forms:   map[string]string{},
	}
}

func (hr *HttpRequest) BuildRequest() (*http.Request, error) {
	parseUrl, err := url.Parse(hr.URL)
	if err != nil {
		return nil, errors.New("url parse err, msg: " + err.Error())
	}

	// append params
	if hr.Queries != nil {
		parseUrl.RawQuery = parseParams(hr.Queries)
	}

	// create body
	var body io.Reader
	if hr.Forms != nil {
		body = strings.NewReader(parseParams(hr.Forms))

		if  _, ok := hr.Headers["Content-Type"]; !ok {
			hr.Headers["Content-Type"] = "application/x-www-form-urlencoded"
		}

	} else if hr.Json != nil {
		body = strings.NewReader(*hr.Json)

		if  _, ok := hr.Headers["Content-Type"]; !ok {
			hr.Headers["Content-Type"] = "application/json"
		}

	} else if hr.Xml != nil {
		body = strings.NewReader(*hr.Xml)

		if  _, ok := hr.Headers["Content-Type"]; !ok {
			hr.Headers["Content-Type"] = "text/xml"
		}

	} else if hr.File != nil {
		body = hr.File

		if  _, ok := hr.Headers["Content-Type"]; !ok {
			hr.Headers["Content-Type"] = "multipart/form-data"
		}

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

func strPrt(s string) *string {
	return &s
}
