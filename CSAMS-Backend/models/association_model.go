package models

import "time"

type AssociationModel struct {
	AssociationID   uint64    `gorm:"primaryKey" json:"association_id"`
	AssociationName string    `gorm:"size:32" json:"association_name"`
	CreateAt        time.Time `json:"create_at"`
	Teacher         string    `gorm:"size:16" json:"teacher"`
	President       string    `gorm:"size:16" json:"president"`
	Introduction    string    `gorm:"size:255" json:"introduction"`
}
