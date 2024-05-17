package global

import (
	"CSAMS-Backend/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	MysqlLog logger.Interface
)
