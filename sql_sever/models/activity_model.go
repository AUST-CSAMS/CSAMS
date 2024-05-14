package models

import "time"

type ActivityModel struct {
	ActivityId          uint      `gorm:"primaryKey" json:"activity_id"`
	ActivityName        string    `gorm:"size:32" json:"activity_name"`
	ActivityTime        time.Time `gorm:"type:datetime" json:"activity_time"`
	ActivityLocation    string    `gorm:"size:32" json:"activity_location"`
	ActivityDescription string    `gorm:"size:256" json:"activity_description"`
	ActivityPoster      string    `gorm:"size:256" json:"activity_poster"`
	ResponsiblePerson   string    `gorm:"size:16" json:"responsible_person"`
	ContactInformation  uint      `json:"contact_information"`
}
