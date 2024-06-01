package global

import (
	"CSAMS-Backend/config"
	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	MysqlLog logger.Interface
	Redis    *redis.Client
	AddrDB   *geoip2.DBReader
)
