package demo

import (
	"fmt"
	. "github.com/971181317/goHttpProMaxPlus"
	"testing"
)

func Test(t *testing.T) {
	request := NewHttpRequest().SetURL("www.baidu.com").
		AppendHeaders(map[string]string{
			"header 1": "headerValue1",
			"header 2": "headerValue1"}).
		AppendCookies(map[string]string{
			"Cookie 1": "cookieValue1",
			"Cookie 2": "cookieValue1"}).
		AppendForms(map[string]string{
			"Form 1": "formValue1",
			"Form 2": "formValue1"}).
		AppendParams(map[string]string{
			"Params 1": "paramValue1",
			"Params 2": "paramValue1",
		})
	fmt.Println(request.BuildRequest())
}
