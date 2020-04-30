package model

import (
	"time"
)

// friend好友列表
type Friend struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Mid uint  //自己
	Oid uint  //对方
}


type FriendDeal struct {
	Id uint
	Nickname string
	Avatar string
}