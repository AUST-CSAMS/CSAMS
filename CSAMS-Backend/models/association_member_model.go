package models

import "time"

type AssociationMemberModel struct {
	AssociationID       uint64    `json:"association_id"`
	StudentID           uint64    `json:"student_id"`
	MemberName          string    `gorm:"size:16" json:"member_name"`
	AssociationPosition string    `gorm:"size:16" json:"association_position"`
	JoiningTime         time.Time `gorm:"type:datetime" json:"joining_time"`
}
