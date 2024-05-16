package models

type ActivityAssignmentLinkModel struct {
	ActivityId    uint          `json:"activity_id"`
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId;references:ActivityId;" json:"activity_model"`
	StudentId     uint          `json:"student_id"`
	AssignmentId  uint          `gorm:"primaryKey" json:"assignment_id"`
}
