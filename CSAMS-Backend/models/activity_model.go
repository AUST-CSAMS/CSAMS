package models

import "time"

type ActivityModel struct {
	ActivityID          uint64    `gorm:"primaryKey" json:"activity_id"`
	ActivityName        string    `gorm:"size:32" json:"activity_name"`
	ActivityStartTime   time.Time `json:"activity_start_time"`
	ActivityEndTime     time.Time `json:"activity_end_time"`
	ActivityLocation    string    `gorm:"size:32" json:"activity_location"`
	ActivityDescription string    `gorm:"type:longtext" json:"activity_description"`
	ActivityPoster      string    `gorm:"size:256" json:"activity_poster"`
	ResponsiblePerson   string    `gorm:"size:16" json:"responsible_person"`
	ContactInformation  uint64    `json:"contact_information"`
	ActivityPoints      uint64    `json:"activity_points"`
}
