package model

import "github.com/jinzhu/gorm"

// cos 操作记录
type Cos struct {
	gorm.Model
	Uid uint
	Name string
	Size int64
	TmpSecretID string `gorm:"size:1000"`
	TmpSecretKey string `gorm:"size:1000"`
	SessionToken string `gorm:"size:1000"`
}
