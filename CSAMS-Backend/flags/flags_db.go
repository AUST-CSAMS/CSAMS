package flags

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models"
	"log"
)

func DB() {
	var err error

	_ = global.DB.SetupJoinTable(&models.AssociationModel{}, "Activities", &models.ActivityAssociationModel{})
	_ = global.DB.SetupJoinTable(&models.UserModel{}, "Activities", &models.ActivityLogModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
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
	log.Println("[ success ] 生成数据库表结构成功！")
}
