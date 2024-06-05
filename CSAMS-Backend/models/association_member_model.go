package models

import "time"

type AssociationMemberModel struct {
	UserID        uint64    `gorm:"primaryKey" json:"-"`
	AssociationID uint64    `json:"association_id"`
	Posts         string    `gorm:"size:16" json:"Posts"`
	JoiningTime   time.Time `json:"joining_time"`
}
