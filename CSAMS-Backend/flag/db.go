package flag

import (
	"sql_sever/global"
	"sql_sever/models"
)

func Makemigrations() {
	var err error
	// 生成表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.AssociationModel{},
			&models.ActivityModel{},
			&models.UserModel{},
			&models.ActivityLogModel{},
			&models.AssociationMemberModel{},
			&models.AssignmentModel{},
			&models.ActivityAssociationModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
