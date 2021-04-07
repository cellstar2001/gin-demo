// package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/ginrpc/api"
	"github.com/xxjwxc/public/mydoc/myswagger"
)

// type ReqTest struct {
// 	Access_token string `json:"access_token"`
// 	UserName     string `json:"user_name" binding:"required"` // 带校验方式
// 	Password     string `json:"password"`
// }

// // Hello ...
// type Hello1 struct {
// }

// // @router /block [get]
// func (s *Hello1) Hello1(c *api.Context, req *ReqTest) {
// 	fmt.Println(req)

// 	c.JSON(http.StatusOK, "ok")
// }

// func main() {
// 	// swagger
// 	myswagger.SetHost("https://localhost:8080")
// 	myswagger.SetBasePath("gmsec")
// 	myswagger.SetSchemes(true, false)
// 	// -----end --
// 	base := ginrpc.New()
// 	router := gin.Default()
// 	// group := router.Group("/xxjwxc")
// 	base.Register(router, new(Hello1)) // 对象注册 like(go-micro)
// 	// router.POST("/test6", base.HandlerFunc(TestFun6))                            // 函数注册
// 	// base.RegisterHandlerFunc(router, []string{"post", "get"}, "/test") // 多种请求方式注册
// 	router.Run(":8080")
// }
