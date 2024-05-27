package models

type ActivityAssociationModel struct {
	ActivityID    uint64           `json:"activity_id"`
	ActivityModel ActivityModel    `gorm:"foreignKey:ActivityID"`
	AssociationID uint64           `json:"association_id"`
	Association   AssociationModel `gorm:"foreignKey:AssociationID"`
}
