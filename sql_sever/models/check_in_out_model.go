package models

type CheckInOutModel struct {
	MODEL
	ActivityId    uint          `json:"activity_id"`
	ActivityModel ActivityModel `gorm:"foreignKey:ActivityId;references:ActivityId;" json:"activity_model"`
	StudentId     uint          `json:"student_id"`
	SignIn        int           `json:"sign_in"`
	SignOut       int           `json:"sign_out"`
}
