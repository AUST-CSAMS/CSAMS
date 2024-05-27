package models

type ActivityAssignmentLinkModel struct {
	ActivityID   uint64 `json:"activity_id"`
	StudentID    uint64 `json:"student_id"`
	AssignmentID uint64 `gorm:"primaryKey" json:"assignment_id"`
}
