package main

import (
	"sql_sever/core"
	"sql_sever/flag"
	"sql_sever/global"
	"sql_sever/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	//表结构转移数据库
	option := flag.Parse()
	//判断是否停止web项目
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//启动gin
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在：%s", addr)
	router.Run(addr)
}
