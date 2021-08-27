package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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
}
