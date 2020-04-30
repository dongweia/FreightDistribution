package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Order 订单模型
type Order struct {
	gorm.Model
	Cid uint  //商品id
	Uid uint //下单用户id
	State string //状态

	Sname       string
	Sphone      string
	Saddress    string
	Sdetail     string
	Rname       string
	Rphone      string
	Raddress    string
	Rdetail     string
	Thingtype   string
	Thingweight int
	Thingmess   string `gorm:"size:1000"`
}

const (
	Dealing string="dealing"
	Finish string="finish"
)

//商家订单列表模型
type OrderBusinessList struct {
	Oid uint
	State string
	Uid uint
	Uname string
	CreatedAt time.Time

	Sname       string
	Sphone      string
	Saddress    string
	Sdetail     string
	Rname       string
	Rphone      string
	Raddress    string
	Rdetail     string
	Thingtype   string
	Thingweight int
	Thingmess   string
}

//客户订单列表模型
type OrderUserList struct {
	Oid uint
	State string
	Mid uint
	Mname string
	CreatedAt time.Time

	Sname       string
	Sphone      string
	Saddress    string
	Sdetail     string
	Rname       string
	Rphone      string
	Raddress    string
	Rdetail     string
	Thingtype   string
	Thingweight int
	Thingmess   string
}