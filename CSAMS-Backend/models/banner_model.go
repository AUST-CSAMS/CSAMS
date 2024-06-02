package models

import (
	"gorm.io/gorm"
	"log"
	"os"
)

type BannerModel struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Path string `json:"path"`                // 图片路径
	Hash string `json:"hash"`                // 图片的hash值，用于判断重复图片
	Name string `gorm:"size:38" json:"name"` // 图片名称
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 本地图片，删除，还要删除本地的存储
	err = os.Remove(b.Path[1:])
	if err != nil {
		log.Print(err)
		return nil
	}
	return nil
}
