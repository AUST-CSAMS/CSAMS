package models

type ReleaseModel struct {
	ActivityModel    ActivityModel    `gorm:"foreignKey:ActivityId" json:"activity_model"`
	AssociationModel AssociationModel `gorm:"foreignKey:AssociationId" json:"association_model"`
}
