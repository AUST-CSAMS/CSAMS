package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id,select($any)" structs:"-"` //主键id
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`           //创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                                 //更新时间
}
