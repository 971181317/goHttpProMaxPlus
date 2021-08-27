# goHttpProMaxPlus

Go quickly request http. 

go快速请求Http

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
