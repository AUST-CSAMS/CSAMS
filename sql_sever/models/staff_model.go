package models

type StaffModel struct {
	StaffID               int       `gorm:"primaryKey" json:"staff_id"`
	UserModel             UserModel `gorm:"foreignKey:StaffID" json:"user_model"`
	College               string    `gorm:"size:32" json:"college"`
	ContactInformation    uint      `json:"contact_information"`
	AffiliatedAssociation string    `gorm:"size:32" json:"affiliated_association"`
}
