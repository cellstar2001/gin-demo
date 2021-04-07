// +build ignore

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 一个简单的http请求  http://localhost:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
