package router

import (
	"gin-demo/controller"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/api").Use(middleware.NeedInit())
	{
		ApiRouter.GET("login", controller.Login) // 创建Api
	}
}
