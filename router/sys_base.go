package router

import (
	"gin-demo/controller"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	{
		BaseRouter.GET("login", controller.Login)
	}
	return BaseRouter
}
