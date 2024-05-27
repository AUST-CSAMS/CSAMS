package models

import "time"

type AssociationMemberModel struct {
	MemberName  UserModel `gorm:"foreignKey:ID" json:"member_name"`
	Posts       string    `gorm:"size:16" json:"Posts"`
	JoiningTime time.Time `json:"joining_time"`
}
