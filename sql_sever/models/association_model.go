package models

import "time"

type AssociationModel struct {
	AssociationId          uint      `gorm:"primaryKey" json:"association_id"`
	AssociationName        string    `gorm:"size:32" json:"name"`
	EstablishmentTime      time.Time `gorm:"type:datetime" json:"establishment_time"`
	ResponsibleTeacherName string    `gorm:"size:16" json:"responsible_teacher_name"`
	PresidentName          string    `gorm:"size:16" json:"president_name"`
	AssociationFunds       int       `json:"association_funds"`
}
