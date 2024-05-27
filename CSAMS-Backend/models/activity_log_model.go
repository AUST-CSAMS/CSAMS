package models

type ActivityLogModel struct {
	ActivityLogID uint64
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityLogID;references:ActivityID"`
	UserLogID     uint64
	UserModel     UserModel `gorm:"foreignKey:UserLogID;references:ID"`
}
