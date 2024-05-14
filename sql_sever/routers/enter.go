package routers

import (
	"github.com/gin-gonic/gin"
	"sql_sever/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 路由分组
	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 路由分层

	// 系统配置api
	routerGroupApp.SettingsRouter()

	return router
}
