package model

import "github.com/jinzhu/gorm"
// Collect 商品收藏
type Collect struct {
	gorm.Model
	Uid uint
	Cid uint
}
