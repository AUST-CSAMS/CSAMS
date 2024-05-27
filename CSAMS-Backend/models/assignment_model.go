package models

type AssignmentModel struct {
	AssignmentID uint64 `json:"assignment_id"`
	Content      string `gorm:"type:longtext" json:"content"`
	Figure1      string `gorm:"size:256" json:"figure_1"`
	Figure2      string `gorm:"size:256" json:"figure_2"`
	Figure3      string `gorm:"size:256" json:"figure_3"`
	Figure4      string `gorm:"size:256" json:"figure_4"`
	Figure5      string `gorm:"size:256" json:"figure_5"`
	Figure6      string `gorm:"size:256" json:"figure_6"`
	Figure7      string `gorm:"size:256" json:"figure_7"`
	Figure8      string `gorm:"size:256" json:"figure_8"`
	Figure9      string `gorm:"size:256" json:"figure_9"`
}
