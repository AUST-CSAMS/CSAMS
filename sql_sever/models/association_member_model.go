package models

import "time"

type AssociationMemberModel struct {
	AssociationId              uint64           `json:"association_id"`
	AssociationModel           AssociationModel `gorm:"many2many:AssociationId" json:"association_model"`
	StudentId                  uint             `json:"student_id"`
	MemberName                 string           `gorm:"size:16" json:"member_name"`
	AssociationPosition        string           `gorm:"size:16" json:"association_position"`
	JoiningTime                time.Time        `gorm:"type:datetime" json:"joining_time"`
	NumberOfActivitiesAttended int              `json:"number_of_activities_attended"`
}
