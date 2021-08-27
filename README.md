# goHttpProMaxPlus

Go quickly request http and Support for AOP

go快速请求Http, 同时支持AOP

> 要不是公司用go，我才不会写他呢！！！

以下英语均为机翻

## Import dependence 导入依赖

```text
go get github.com/971181317/goHttpProMaxPlus
```

## Quick start 快速开始

First build a server with the `gin` framework

首先用`gin`框架搭建一个服务器

```go
import "github.com/gin-gonic/gin"

r := gin.Default()
r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "msg": "This is a GET method",
    })
})

r.POST("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "msg": "This is a POST method",
    })
})

r.Run() // listen and serve on 0.0.0.0:8080
```

Next use `goHttpProMaxPlus` to quickly request `Get` and `Post`

接下来使用`goHttpProMaxPlus`快速请求`Get`和`Post`

```go
import . "github.com/971181317/goHttpProMaxPlus"

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
```

result

结果

```text
{"msg":"This is a GET method"}
{"msg":"This is a POST method"}
```

It's pretty easy, right?

芜湖，起飞（老师，老师，我已经会get和post了）！！！

But what if you want to use `Header`, `Cookie`, or `Form`?  

但如果想使用`Header`，`Cookie`或者`Form`怎么办呢？

```go
// cookie, header, form: map[string]string
// data: *io.Reader
GetDefaultClient().GetWithCookieAndHeader(url, cookie, header)
GetDefaultClient().PostWithForm(url, form)
GetDefaultClient().PostWithCookieHeaderAndForm(url, cookie, header, form)
GetDefaultClient().PostWithIoData(url sting, data)
GetDefaultClient().PostWithCookieHeaderAndIoData(url, cookie, header, data)
```

## AOP

**5 Aspects:**

1. BeforeClientBuild
2. AfterClientBuild
3. BeforeRequestBuild 
4. AfterRequestBuild
5. AfterResponseCreate

## Custom HTTP request 自定义请求

**4 Steps：**

1. create client

```go
// AspectModel 切片模组
type AspectModel func(...interface{})

// HttpClient A client send request
type HttpClient struct {
	c                   *http.Client
	BeforeClientBuild   AspectModel
	AfterClientBuild    AspectModel
	BeforeRequestBuild  AspectModel
	AfterRequestBuild   AspectModel
	AfterResponseCreate AspectModel
	AspectArgs          []interface{} // Only the most recent assignment will be kept
}
```

create it and use aspect

```go
// This method execute BeforeClientBuild and AfterClientBuild.
func NewClientX(client *http.Client,
	beforeClientBuild, afterClientBuild, beforeRequestBuild, afterRequestBuild, afterResponseCreate AspectModel,
	args ...interface{}) *HttpClient
```

2. create request

```go
// Forms > Json > Xml > File > ReaderBody
type HttpRequest struct {
	Method     HttpMethod
	URL        string
	Cookies    map[string]string
	Headers    map[string]string
	Queries    map[string]string
	Forms      map[string]string
	Json       *string
	Xml        *string
	File       *os.File
	ReaderBody *io.Reader
}
```

Two types of assignment

```go
// 1. chain
req := NewHttpRequest().
		SetMethod(POST).
		SetURL("http://localhost:8080").
		AppendHeader("header", "headerValue").
		AppendCookie("cookie", "cookieValue").
		AppendForm("form", "formValue").
		AppendQuery("query", "queryValue")

// 2.
req := &HttpRequest{
    Method     HttpMethod
	URL        "http://localhost:8080"
	Cookies    ...
	Headers    ...
	Queries    ...
	Forms      ...
    ...
}
```

3. do request

```go
resp, err := client.Do(req)
```

4. get request body

```go
type HttpResponse struct {
	resp    *http.Response
	Headers map[string]string
	body    io.ReadCloser
	bodyStr *string
}

func (hr *HttpResponse) ParseJson(v interface{})
func (hr *HttpResponse) GetRespStr() string
```

