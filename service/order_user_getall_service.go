package service
import (
"FreightDistribution/logger"
"FreightDistribution/model"
"FreightDistribution/serializer"
"github.com/gin-gonic/gin"
)

type GetAllOrderUserService struct {
	Type string `form:"type" json:"type" binding:"required"`
}


func (service *GetAllOrderUserService)GetAll(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言

	list:=[]model.OrderUserList{}
	if err:=model.DB.Table("orders").Select("orders.id as oid,orders.created_at as created_at,orders.state as state,commodities.mid as mid,users.nickname as mname,orders.sname as sname,orders.sphone as sphone,orders.saddress as saddress,orders.sdetail as sdetail,orders.rname as rname,orders.rphone as rphone,orders.raddress as raddress,orders.rdetail as rdetail,orders.thingtype as thingtype,orders.thingweight as thingweight,orders.thingmess as thingmess").Joins("left join commodities on orders.cid = commodities.id").Joins("left join users on commodities.mid = users.id").Where("orders.uid=? and state=?",u.ID,service.Type).Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	//if err:=model.DB.Table("commodities").Select("orders.id as oid,orders.created_at as created_at,orders.state as state,commodities.mid as mid,users.nickname as mname,orders.sname as sname,orders.sphone as sphone,orders.saddress as saddress,orders.sdetail as sdetail,orders.rname as rname,orders.rphone as rphone,orders.raddress as raddress,orders.rdetail as rdetail,orders.thingtype as thingtype,orders.thingweight as thingweight,orders.thingmess as thingmess").Joins("left join orders on orders.cid = commodities.id").Joins("left join users on commodities.mid = users.id").Where("orders.uid=? and orders.state=?",u.ID,service.Type).Scan(&list).Error;err!=nil{
	//	logger.Log().Error("数据库查询失败", err)
	//	return  serializer.DBErr("数据库操作失败",err)
	//}
	return serializer.BuildOrderUserListResponse(&list)
}

