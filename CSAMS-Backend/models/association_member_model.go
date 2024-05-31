package models

import "time"

type AssociationMemberModel struct {
	ID          uint64      `json:"id"`
	Users       []UserModel `gorm:"foreignKey:AssociationID" json:"users"`
	Posts       string      `gorm:"size:16" json:"Posts"`
	JoiningTime time.Time   `json:"joining_time"`
}
