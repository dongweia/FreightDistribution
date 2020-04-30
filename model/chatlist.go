package model

import (
	"time"
)

// chatlist 聊天列表
type Chatlist struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Sid uint
	Rid uint
	LastContent string `gorm:"size:5000"`
	Type uint   //0普通消息 1图片

}

//组合查询

// chatlist 聊天列表
type GetChatlist struct {
	Id uint
	CreatedAt time.Time
	LastContent string
	Type uint  //0普通消息 1图片
	Nickname  string
	Avatar    string
}