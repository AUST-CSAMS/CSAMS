package models

type UserModel struct {
	MODEL
	StudentStaffId uint   `gorm:"primaryKey" json:"student_staff_id"`
	Password       string `gorm:"type:varchar(128)" json:"password"`
	Name           string `gorm:"type:varchar(16)" json:"name"`
	Age            int    `gorm:"type:tinyint(3)" json:"age"`
	Gender         string `gorm:"type:varchar(3)" json:"gender"`
	ProfilePicture string `gorm:"type:varchar(256)" json:"profile_picture"`
}
