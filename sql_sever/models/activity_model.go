package models

type ActivityModel struct {
	MODEL
	ActivityId          uint   `gorm:"primaryKey" json:"activity_id"`
	ActivityName        string `gorm:"size:32" json:"activity_name"`
	ActivityLocation    string `gorm:"size:32" json:"activity_location"`
	ActivityDescription string `gorm:"size:256" json:"activity_description"`
	ActivityPoster      string `gorm:"size:256" json:"activity_poster"`
	ResponsiblePerson   string `gorm:"size:16" json:"responsible_person"`
	ContactInformation  uint   `json:"contact_information"`
}
