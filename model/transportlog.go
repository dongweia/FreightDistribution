package model

import "github.com/jinzhu/gorm"

// Transport 运输记录模型
type Transportlog struct {
	gorm.Model
	Oid  uint
	Content string
	Phone string
}

