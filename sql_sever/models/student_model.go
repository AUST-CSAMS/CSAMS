package models

type StudentModel struct {
	MODEL
	StudentId             string    `gorm:"primaryKey" json:"student_id"`
	UserModel             UserModel `gorm:"foreignKey:StudentStaffId;references:StudentId;" json:"user_model"`
	College               string    `gorm:"type:varchar(32)" json:"college"`
	Major                 string    `gorm:"type:varchar(16)" json:"major"`
	ContactInformation    uint      `json:"contact_information"`
	AffiliatedAssociation string    `gorm:"type:varchar(32)" json:"affiliated_association"`
	IntegrityScore        int       `gorm:"type:tinyint(3)" json:"integrity_score"`
}
