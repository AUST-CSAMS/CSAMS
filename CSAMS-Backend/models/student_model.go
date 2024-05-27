package models

type StudentModel struct {
	StudentID                  uint   `gorm:"primaryKey" json:"student_id"`
	College                    string `gorm:"size:32" json:"college"`
	Major                      string `gorm:"size:32" json:"major"`
	ContactInformation         uint   `json:"contact_information"`
	AffiliatedAssociation      string `gorm:"size:32" json:"affiliated_association"`
	NumberOfActivitiesAttended int    `json:"number_of_activities_attended"`
	IntegrityScore             int    `json:"integrity_score"`
	ActivityPoints             uint   `json:"activity_points"`
}
