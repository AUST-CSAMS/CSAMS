package models

import "CSAMS-Backend/models/ctype"

type UserModel struct {
	ID                uint64                 `gorm:"primaryKey" json:"id"`
	Password          string                 `gorm:"size:128" json:"password"`
	Name              string                 `gorm:"size:16" json:"name"`
	Age               int                    `gorm:"size:3" json:"age"`
	Gender            string                 `gorm:"size:3" json:"gender"`
	Avatar            string                 `gorm:"size:256" json:"avatar"`
	Role              ctype.Role             `gorm:"size:3" json:"role"`
	Major             ctype.Major            `gorm:"size:5" json:"major"`
	Tel               uint64                 `json:"tel"`
	Integrity         int                    `json:"integrity"`
	Score             uint64                 `json:"score"`
	AssociationID     *uint64                `json:"association_id"`
	AssociationMember AssociationMemberModel `gorm:"foreignKey:AssociationID" json:"-"`
	Assignments       []AssignmentModel      `gorm:"foreignKey:UserID" json:"-"`
	Activities        []ActivityModel        `gorm:"many2many:activity_log_models;joinForeignKey:UserID;JoinReferences:ActivityID" json:"-"`
}
