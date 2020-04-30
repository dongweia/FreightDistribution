package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//用户聊天列表 添加一个聊天
type AddFriendService struct {
	Oid uint `form:"oid" json:"oid" binding:"required"`
}

//添加
func (service *AddFriendService)Add(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	//验证是否重复
	count:=0
	if err:=model.DB.Model(&model.Friend{}).Where(&model.Friend{Mid:u.ID, Oid:service.Oid,}).Count(&count).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	if count>0{
		return serializer.BuildFriendResponse(1) //1代表已存在
	}
	friend:=model.Friend{
		Mid:u.ID,
		Oid:service.Oid,
	}
	//创建
	if err:=model.DB.Create(&friend).Error;err!=nil{
		//添加日志
		logger.Log().Error("数据库创建用户失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildFriendResponse(0)
}


