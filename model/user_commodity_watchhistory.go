package model

import "github.com/jinzhu/gorm"

//History 商品浏览记录
type WatchHistory struct {
	gorm.Model
	Uid uint
	Cid uint
}

