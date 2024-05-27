package models

type ActivityLogModel struct {
	ActivityID uint64        `json:"activity_id"`
	Activity   ActivityModel `gorm:"foreignKey:ActivityID"`
	UserID     uint64        `json:"student_id"`
	User       UserModel     `gorm:"foreignKey:ID"`
}
