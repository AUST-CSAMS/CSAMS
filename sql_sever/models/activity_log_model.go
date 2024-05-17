package models

type ActivityLogModel struct {
	ActivityId    uint          `json:"activity_id"`
	ActivityModel ActivityModel `gorm:"many2many:ActivityId" json:"activity_model"`
	StudentId     string        `json:"student_id"`
	StudentModel  StudentModel  `gorm:"many2many:StudentId" json:"student_model"`
}
