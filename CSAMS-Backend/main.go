package main

import (
	"CSAMS-Backend/core"
	"CSAMS-Backend/flags"
	"CSAMS-Backend/global"
	"CSAMS-Backend/routers"
	"CSAMS-Backend/utils"
	"log"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	// 命令行参数绑定
	option := flags.Parse()
	if option.Run() {
		return
	}
	//初始化路由
	router := routers.InitRouter()
	//返回ip地址
	addr := global.Config.System.Addr()
	//打印本机所有可以ip
	utils.PrintSystem()
	//运行
	err := router.Run(addr)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
