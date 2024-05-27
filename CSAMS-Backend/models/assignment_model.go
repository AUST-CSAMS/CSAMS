package models

type AssignmentModel struct {
	AssignmentID uint64    `gorm:"primaryKey" json:"assignment_id"`
	Content      string    `json:"content"`
	UserID       uint64    `json:"user_id"`
	UserModel    UserModel `gorm:"foreignKey:UserID;references:ID"`
}
