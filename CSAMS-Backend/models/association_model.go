package models

import "time"

type AssociationModel struct {
	ID              uint64          `gorm:"primaryKey" json:"id"`
	AssociationName string          `gorm:"size:32" json:"association_name"`
	CreateAt        time.Time       `json:"create_at"`
	Teacher         string          `gorm:"size:16" json:"teacher"`
	President       string          `gorm:"size:16" json:"president"`
	Introduction    string          `gorm:"size:256" json:"introduction"`
	Activities      []ActivityModel `gorm:"many2many:activity_association_models;joinForeignKey:AssociationID;JoinReferences:ActivityID" json:"activities"`
}
