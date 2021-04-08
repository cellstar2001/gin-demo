package initialize

import (
	"gin-demo/global"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.GVA_LOG.Info("use middleware logger")
	// Router.Use(middleware.Cors()) // 如需跨域可以打开
	// global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// PublicGroup := Router.Group("api")
	// {
	// 	router.InitApiRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	// 	router.InitApiRouter(PublicGroup) // 自动初始化相关
	// }
	global.GVA_LOG.Info("router register success")
	return Router
}
