package goHttpProMaxPlus

import "net/http"

type HttpRequest struct {
	request *http.Request
}

func (hr *HttpRequest) AppendCookie(name, value string) *HttpRequest {
	hr.request.AddCookie(&http.Cookie{
		Name:  name,
		Value: value,
	})
	return hr
}

func (hr *HttpRequest) AppendCookies(cookies map[string]string) *HttpRequest {
	for k, v := range cookies {
		hr.AppendCookie(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetCookie(name string) string {
	cookie, err := hr.request.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (hr *HttpRequest) ClearAllCookie() *HttpRequest {
	hr.request.Header.Del("Cookie")
	return hr
}

func (hr *HttpRequest) ReplaceCookie(cookies map[string]string) {
	hr.ClearAllCookie()
	for k, v := range cookies {
		hr.AppendCookie(k, v)
	}
}

func (hr *HttpRequest) AppendHeader(name, value string) *HttpRequest {
	hr.request.Header.Set(name, value)
	return hr
}

func (hr *HttpRequest) AppendHeaders(headers map[string]string) *HttpRequest {
	for k, v := range headers {
		hr.AppendHeader(k, v)
	}
	return hr
}

func (hr *HttpRequest) GetHeader(name string) string {
	return hr.request.Header.Get(name)
}

func (hr *HttpRequest) ClearAllHeader() *HttpRequest {
	hr.request.Header = http.Header{}
	return hr
}

func (hr *HttpRequest) ReplaceHeader(header map[string]string) *HttpRequest {
	hr.ClearAllHeader()
	hr.AppendHeaders(header)
	return hr
}

func (hr *HttpRequest) GetGoRequest(header map[string]string) *http.Request {
	return hr.request
}
