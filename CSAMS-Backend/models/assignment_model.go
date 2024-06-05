package models

type AssignmentModel struct {
	ID         uint64        `json:"id"`
	Content    string        `json:"content"`
	ActivityID uint64        `json:"activity_id"`
	Activity   ActivityModel `gorm:"foreignKey:ActivityID" json:"-"`
	UserID     uint64        `json:"user_id"`
	User       UserModel     `gorm:"foreignKey:UserID" json:"-"`
	IsSubmit   bool          `json:"is_submit"`
	IsFinish   bool          `json:"is_finish"`
	IsCorrect  bool          `json:"is_correct"`
}
