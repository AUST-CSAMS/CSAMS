package core

import (
	"CSAMS-Backend/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql，取消gorm连接")
		return nil
	}

	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	//通过gorm.Open()函数连接到MySQL数据库，并设置日志记录器为上一步配置的mysqlLogger
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 最多可容纳
	sqlDB.SetMaxOpenConns(100)
	// 连接最大复用时间，不能超过mysql的wait_timeout
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
