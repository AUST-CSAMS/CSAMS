package models

type ReleaseModel struct {
	ActivityId    uint `gorm:"-" json:"activity_id"`
	AssociationId uint `gorm:"-" json:"association_id"`
}
