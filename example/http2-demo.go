// +build ignore

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("getDemo", getting)
}

func getting() {
	fmt.Println("11111111111")
}
