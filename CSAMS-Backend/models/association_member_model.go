package models

import "time"

type AssociationMemberModel struct {
	ID          uint64      `gorm:"primaryKey" json:"id"`
	Users       []UserModel `gorm:"foreignKey:AssociationID" json:"-"`
	Posts       string      `gorm:"size:16" json:"Posts"`
	JoiningTime time.Time   `json:"joining_time"`
}
