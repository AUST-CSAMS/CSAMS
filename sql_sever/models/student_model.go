package models

type StudentModel struct {
	MODEL
	StudentId             string    `gorm:"primaryKey" json:"student_id"`
	UserModel             UserModel `gorm:"foreignKey:StudentStaffId" json:"user_model"`
	College               string    `gorm:"size:32" json:"college"`
	Major                 string    `gorm:"size:16" json:"major"`
	ContactInformation    uint      `json:"contact_information"`
	AffiliatedAssociation string    `gorm:"size:32" json:"affiliated_association"`
	IntegrityScore        int       `gorm:"size:3" json:"integrity_score"`
}
