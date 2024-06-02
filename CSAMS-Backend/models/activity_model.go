package models

import "time"

type ActivityModel struct {
	ID                  uint64                     `gorm:"primaryKey" json:"id"`
	ActivityName        string                     `gorm:"size:32" json:"activity_name"`
	StartTime           time.Time                  `json:"start_time"`
	EndTime             time.Time                  `json:"end_time"`
	Location            string                     `gorm:"size:32" json:"location"`
	Introduction        string                     `gorm:"size:255" json:"introduction"`
	BannerID            uint64                     `json:"banner_id"`
	BannerModel         BannerModel                `gorm:"foreignKey:BannerID" json:"-"`
	ResponsiblePerson   string                     `gorm:"size:16" json:"responsible_person"`
	Tel                 uint64                     `json:"tel"`
	Score               uint64                     `json:"score"`
	Assignments         []AssignmentModel          `gorm:"foreignKey:ActivityID" json:"-"`
	ActivityAssociation []ActivityAssociationModel `gorm:"foreignKey:ActivityID" json:"-"`
	ActivityLog         []ActivityLogModel         `gorm:"foreignKey:ActivityID" json:"-"`
}
