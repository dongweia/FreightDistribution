package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

type UserOrderEnterService struct {
	Oid uint `form:"oid" json:"oid" binding:"required"`
}
func (service *UserOrderEnterService)Enter(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言


	if err:=model.DB.Model(&model.Order{}).Where("id=? and uid=?",service.Oid,u.ID).Update("state",model.Finish).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildOrderEnterResponse()
}