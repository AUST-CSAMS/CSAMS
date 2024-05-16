package models

type ActivityModel struct {
	MODEL
	ActivityId          uint   `gorm:"primaryKey" json:"activity_id"`
	ActivityName        string `gorm:"type:varchar(32)" json:"activity_name"`
	ActivityLocation    string `gorm:"type:varchar(32)" json:"activity_location"`
	ActivityDescription string `gorm:"type:varchar(256)" json:"activity_description"`
	ActivityPoster      string `gorm:"type:varchar(256)" json:"activity_poster"`
	ResponsiblePerson   string `gorm:"type:varchar(16)" json:"responsible_person"`
	ContactInformation  uint   `json:"contact_information"`
}
