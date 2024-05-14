package models

type ActivityAssignmentLinkModel struct {
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId" json:"activity_model"`
	StudentId     uint          `json:"student_id"`
	AssignmentId  uint          `gorm:"primaryKey" json:"assignment_id"`
}
