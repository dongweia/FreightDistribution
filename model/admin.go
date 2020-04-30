package model

import "github.com/jinzhu/gorm"

// Admin 管理员模型
type Admin struct {
	gorm.Model
	Uid uint
	Lv int
}
