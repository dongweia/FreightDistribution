package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// chatlist 聊天信息列表
type Chatmess struct {
	gorm.Model
	Sid uint
	Rid uint
	Content string `gorm:"size:5000"`
	Status uint    //0未读 1已读
	Type uint   //0普通消息 1图片
}


//聊天历史记录加载
type ChatHistory struct {
	Sid uint
	Sname string
	Savatar string
	Content string
	Type uint
	Time time.Time
}
