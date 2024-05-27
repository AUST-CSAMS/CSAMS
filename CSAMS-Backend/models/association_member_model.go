package models

import "time"

type AssociationMemberModel struct {
	MemberID    uint      `json:"member_id"`
	Usermodel   UserModel `gorm:"foreignKey:MemberID;AssociationForeignKey:ID" json:"member_name"`
	Posts       string    `gorm:"size:16" json:"Posts"`
	JoiningTime time.Time `json:"joining_time"`
}
