package serializer

import "FreightDistribution/model"

// OrderUserList //用户查看订单
type OrderUserList struct {
	Oid uint `json:"oid"`
	State string `json:"state"`
	Mid uint `json:"mid"`
	Mname string `json:"mname"`
	Time int64 `json:"time"`

	Sname       string `json:"sname"`
	Sphone      string `json:"sphone"`
	Saddress    string `json:"saddress"`
	Sdetail     string `json:"sdetail" `
	Rname       string `json:"rname"`
	Rphone      string `json:"rphone"`
	Raddress    string `json:"raddress" `
	Rdetail     string `json:"rdetail" `
	Thingtype   string `json:"thingtype" `
	Thingweight int    `json:"thingweight" `
	Thingmess   string `json:"thingmess" `

}


//
type OrderBusinessList struct {
	Oid uint `json:"oid"`
	State string `json:"state"`
	Uid uint `json:"uid"`
	Uname string `json:"uname"`
	Time int64 `json:"time"`

	Sname       string `json:"sname"`
	Sphone      string `json:"sphone"`
	Saddress    string `json:"saddress"`
	Sdetail     string `json:"sdetail" `
	Rname       string `json:"rname"`
	Rphone      string `json:"rphone"`
	Raddress    string `json:"raddress" `
	Rdetail     string `json:"rdetail" `
	Thingtype   string `json:"thingtype" `
	Thingweight int    `json:"thingweight" `
	Thingmess   string `json:"thingmess" `
}

// BuildOrderUserList 序列化收藏商品列表
func BuildOrderUserList(orderlist *[]model.OrderUserList) *[]OrderUserList {
	newuserlist:=make([]OrderUserList,len(*orderlist))
	for i,v:=range *orderlist{
		newuserlist[i].Oid=v.Oid
		newuserlist[i].State=v.State
		newuserlist[i].Mid=v.Mid
		newuserlist[i].Mname=v.Mname
		newuserlist[i].Time=v.CreatedAt.Unix()

		newuserlist[i].Sname=v.Sname
		newuserlist[i].Sphone=v.Sphone
		newuserlist[i].Saddress=v.Saddress
		newuserlist[i].Sdetail=v.Sdetail

		newuserlist[i].Rname=v.Rname
		newuserlist[i].Rphone=v.Rphone
		newuserlist[i].Raddress=v.Raddress
		newuserlist[i].Rdetail=v.Rdetail

		newuserlist[i].Thingtype=v.Thingtype
		newuserlist[i].Thingweight=v.Thingweight
		newuserlist[i].Thingmess =v.Thingmess

	}
	return &newuserlist
}

// BuildOrderBusinessList 序列化收藏商品列表
func BuildOrderBusinessList(orderlist *[]model.OrderBusinessList) *[]OrderBusinessList {
	newbusinesslist:=make([]OrderBusinessList,len(*orderlist))
	for i,v:=range *orderlist{
		newbusinesslist[i].Oid=v.Oid
		newbusinesslist[i].State=v.State
		newbusinesslist[i].Uid=v.Uid
		newbusinesslist[i].Uname=v.Uname
		newbusinesslist[i].Time=v.CreatedAt.Unix()

		newbusinesslist[i].Sname=v.Sname
		newbusinesslist[i].Sphone=v.Sphone
		newbusinesslist[i].Saddress=v.Saddress
		newbusinesslist[i].Sdetail=v.Sdetail

		newbusinesslist[i].Rname=v.Rname
		newbusinesslist[i].Rphone=v.Rphone
		newbusinesslist[i].Raddress=v.Raddress
		newbusinesslist[i].Rdetail=v.Rdetail

		newbusinesslist[i].Thingtype=v.Thingtype
		newbusinesslist[i].Thingweight=v.Thingweight
		newbusinesslist[i].Thingmess =v.Thingmess

	}
	return &newbusinesslist
}


// BuildOrderEnterResponse 序列化上传成功响应
func BuildOrderEnterResponse() *Response {
	return &Response{
	}
}

// BuildOrderAddResponse 序列化上传成功响应
func BuildOrderAddResponse(oid uint) *Response {
	return &Response{
		Data:oid,
	}
}

//BuildOrderUserListResponse //序列化用户列表响应
func BuildOrderUserListResponse(orderlist *[]model.OrderUserList) *Response {
	return &Response{
		Data:BuildOrderUserList(orderlist),
	}
}

//BuildOrderBusinessListResponse //序列化商家列表响应
func BuildOrderBusinessListResponse(orderlist *[]model.OrderBusinessList) *Response {
	return &Response{
		Data:BuildOrderBusinessList(orderlist),
	}
}