package main

import (
	"CSAMS-Backend/core"
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
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
	/*
		option := flags.Parse()
		if option.Run() {
			return
		}
	*/
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.ActivityAssociationModel{},
			&models.ActivityLogModel{},
			&models.AssignmentModel{},
			&models.ActivityModel{},
			&models.UserModel{},
			&models.AssociationModel{},
			&models.AssociationMemberModel{},
		)
	if err != nil {
		log.Fatalf("[ error ] 生成数据库表结构失败")
		return
	}
	//初始化路由
	router := routers.InitRouter()
	//返回ip地址
	addr := global.Config.System.Addr()
	//打印本机所有可以ip
	utils.PrintSystem()
	//运行
	err = router.Run(addr)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
