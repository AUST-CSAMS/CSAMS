package models

type StaffModel struct {
	StaffID               uint64 `gorm:"primaryKey" json:"staff_id"`
	College               string `gorm:"size:32" json:"college"`
	ContactInformation    uint64 `json:"contact_information"`
	AffiliatedAssociation string `gorm:"size:32" json:"affiliated_association"`
}
