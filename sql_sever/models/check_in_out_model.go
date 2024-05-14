package models

type CheckInOutModel struct {
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId" json:"activity_model"`
	StudentId     uint          `json:"student_id"`
	SignIn        int           `json:"sign_in"`
	SignOut       int           `json:"sign_out"`
}
