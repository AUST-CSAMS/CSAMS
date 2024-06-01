package models

// LoginDataModel 统计用户登录数据
type LoginDataModel struct {
	UserID    uint64    `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` // 登录ip
	Name      string    `gorm:"size:42" json:"name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` // 登录设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
