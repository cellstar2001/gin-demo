package router

import (
	"gin-demo/controller"

	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.GET("createApi", controller.Login) // 创建Api
	}
}
