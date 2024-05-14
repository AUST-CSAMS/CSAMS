package models

type ActivityLogModel struct {
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId" json:"activity_model"`
	StudentModel  StudentModel  `gorm:"foreignKey:StudentId" json:"student_model"`
}
