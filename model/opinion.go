package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Opinion 商品收藏
type Opinion struct {
	gorm.Model
	Uid uint
	Content string `gorm:"size:1000"`
}



type OpinionUser struct {
	Uid uint
	Uname string
	Uavatar string
	Content string
	CreatedAt time.Time
}