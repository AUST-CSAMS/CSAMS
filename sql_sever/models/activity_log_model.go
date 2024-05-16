package models

type ActivityLogModel struct {
	ActivityId    uint          `json:"activity_id"`
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId;references:ActivityId;" json:"activity_model"`
	StudentId     string        `json:"student_id"`
	StudentModel  StudentModel  `gorm:"foreignKey:StudentId;references:StudentId;" json:"student_model"`
}
