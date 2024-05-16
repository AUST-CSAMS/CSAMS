package routers

import (
	"CSAMS-Backend/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	//设置gin模式
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//将指定目录下的文件提供给客户端
	//"uploads" 是URL路径前缀，http.Dir("uploads")是实际文件系统中存储文件的目录
	router.StaticFS("uploads", http.Dir("uploads"))
	//注册swagger路由
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//创建路由组
	/*
		apiRouterGroup := router.Group("api")
		routerGroupApp := RouterGroup{apiRouterGroup}
	*/
	// 系统配置api
	return router
}
