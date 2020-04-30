package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//添加一个收藏
type AddOrderService struct {
	Cid uint `form:"cid" json:"cid" binding:"required"`
	Sname string `form:"sname" json:"sname" binding:"required"`
	Sphone string `form:"sphone" json:"sphone" binding:"required"`
	Saddress string `form:"saddress" json:"saddress" binding:"required"`
	Sdetail string `form:"sdetail" json:"sdetail" `
	Rname string `form:"rname" json:"rname" binding:"required"`
	Rphone string `form:"rphone" json:"rphone" binding:"required"`
	Raddress string `form:"raddress" json:"raddress" binding:"required"`
	Rdetail string `form:"rdetail" json:"rdetail" `
	Thingtype string `form:"thingtype" json:"thingtype" binding:"required"`
	Thingweight int `form:"thingweight" json:"thingweight" binding:"required"`
	ThingMess string `form:"thingmess" json:"thingmess" `
}

//添加
func (service *AddOrderService)Add(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言

	order:=&model.Order{
		Cid:         service.Cid,
		Uid:         u.ID,
		State:       model.Dealing,
		Sname:       service.Sname,
		Sphone:      service.Sphone,
		Saddress:    service.Saddress,
		Sdetail:     service.Sdetail,
		Rname:       service.Rname,
		Rphone:      service.Rphone,
		Raddress:    service.Raddress,
		Rdetail:     service.Rdetail,
		Thingtype:   service.Thingtype,
		Thingweight: service.Thingweight,
		Thingmess:   service.ThingMess,
	}

	//创建
	if err:=model.DB.Create(&order).Error;err!=nil{
		//添加日志
		logger.Log().Error("数据库创建订单失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildOrderAddResponse(order.ID)
}



