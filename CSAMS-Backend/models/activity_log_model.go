package models

type ActivityLogModel struct {
	ActivityID uint64 `json:"activity_id"`
	StudentID  uint64 `json:"student_id"`
}
