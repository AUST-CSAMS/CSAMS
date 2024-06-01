package routers

import (
	"CSAMS-Backend/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// 设置全局变量store
var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	//使用session中间件
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("register", app.UserRegisterView)
}
