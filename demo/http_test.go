package demo_test

import (
	"fmt"
	. "github.com/971181317/goHttpProMaxPlus"
	"testing"
)

func TestQuickStart(t *testing.T) {
	resp, err := GetDefaultClient().Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetRespStr())

	resp, err = GetDefaultClient().Post("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetRespStr())
}

func Test(t *testing.T) {
	req := NewHttpRequest().
		SetMethod(POST).
		SetURL("http://localhost:8080").
		AppendHeader("header", "headerValue").
		AppendCookie("cookie", "cookieValue").
		AppendForm("form", "formValue").
		AppendQuery("query", "queryValue")

	resp, _ := GetDefaultClient().Do(req)

	fmt.Println(resp.GetRespStr())
}
