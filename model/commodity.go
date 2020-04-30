package model

import "github.com/jinzhu/gorm"

// Commodity  商品模型
type Commodity struct {
	gorm.Model
	Mid uint
	Title string
	Describe string `gorm:"size:5000"`
	Image string `gorm:"size:1000"`
	Lable string
	Address string `gorm:"size:2000"`
}
