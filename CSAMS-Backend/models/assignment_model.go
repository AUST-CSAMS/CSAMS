package models

type AssignmentModel struct {
	ID         uint64        `json:"assignment_id"`
	Content    string        `json:"content"`
	ActivityID uint64        `json:"activity_id"`
	Activity   ActivityModel `gorm:"foreignKey:ActivityID" json:"-"`
	UserID     uint64        `json:"user_id"`
	User       UserModel     `gorm:"foreignKey:UserID" json:"-"`
	IsCorrect  bool          `json:"is_correct"`
}
