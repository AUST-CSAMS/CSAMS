package models

import "time"

type AssociationMemberModel struct {
	UserID        uint64    `gorm:"primaryKey" json:"user_id"`
	AssociationID uint64    `json:"association_id"`
	Posts         string    `gorm:"size:16" json:"posts"`
	JoiningTime   time.Time `json:"joining_time"`
}
