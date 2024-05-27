package models

type UserModel struct {
	StudentStaffId uint   `gorm:"primaryKey" json:"student_staff_id"`
	Password       string `gorm:"size:128" json:"password"`
	Name           string `gorm:"size:16" json:"name"`
	Age            int    `gorm:"size:3" json:"age"`
	Gender         string `gorm:"size:3" json:"gender"`
	ProfilePicture string `gorm:"size:256" json:"profile_picture"`
}
