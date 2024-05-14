package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sql_sever/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
