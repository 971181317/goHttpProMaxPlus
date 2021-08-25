package demo

import (
	"fmt"
	. "github.com/971181317/goHttpProMaxPlus"
	"net/url"
	"testing"
)

func Test(t *testing.T) {
	parse, _ := url.Parse("https://www.baidu.com/indx?ad=fsji")
	hr := GetHttpRequest(GET, parse)
	hr.AppendHeaders(map[string]string{
		"header1": "1",
		"header2": "1",
	}).AppendCookies(map[string]string{
		"Cookie1": "1",
		"Cookie2": "1",
	}).AppendForms(map[string]string{
		"Form1": "1",
		"Form2": "1",
	}).AppendParams(map[string]string{
		"1Params水电费": "1",
		"2Params发送方": "1",
	})
	fmt.Println(hr.GetCookie("Cookie1"))
	fmt.Println(hr.GetHeader("header1"))
	fmt.Println(hr.GetForm("Form1"))
	fmt.Println(hr.GetParam("1Params水电费"))
	request, _ := BuildRequest(hr)
	fmt.Println(request)
}
