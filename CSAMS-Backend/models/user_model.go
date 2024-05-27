package models

type UserModel struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Password      string `gorm:"size:128" json:"password"`
	Name          string `gorm:"size:16" json:"name"`
	Age           int    `gorm:"size:3" json:"age"`
	Gender        string `gorm:"size:3" json:"gender"`
	Avatar        string `gorm:"size:256" json:"avatar"`
	Role          int    `gorm:"size:3" json:"role"`
	Major         string `gorm:"size:32" json:"major"`
	Tel           uint   `json:"contact_information"`
	Integrity     int    `json:"integrity"`
	ActivityScore uint   `json:"activity_score"`
}
