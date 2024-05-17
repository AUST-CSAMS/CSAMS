package models

type ActivityAssignmentLinkModel struct {
	AssignmentId  uint          `gorm:"primaryKey" json:"assignment_id"`
	ActivityId    uint          `json:"activity_id"`
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId" json:"activity_model"`
	StudentId     uint          `json:"student_id"`
	StudentModel  StudentModel  `gorm:"foreignKey:StudentId" json:"student_model"`
}
