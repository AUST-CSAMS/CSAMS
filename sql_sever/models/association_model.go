package models

import "time"

type AssociationModel struct {
	AssociationId          uint64    `gorm:"primaryKey" json:"association_id"`
	AssociationName        string    `gorm:"type:varchar(32)" json:"name"`
	EstablishmentTime      time.Time `gorm:"type:datetime" json:"establishment_time"`
	ResponsibleTeacherName string    `gorm:"type:varchar(16)" json:"responsible_teacher_name"`
	PresidentName          string    `gorm:"type:varchar(16)" json:"president_name"`
	AssociationFunds       int       `json:"association_funds"`
}
