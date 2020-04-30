package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//用户聊天列表 添加一个聊天
type AddChatService struct {
	Rid uint `form:"rid" json:"rid" binding:"required"`
}

//添加
func (service *AddChatService)Add(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	//验证是否重复
	count:=0
	if err:=model.DB.Model(&model.Chatlist{}).Where(&model.Chatlist{Sid:u.ID, Rid:service.Rid,}).Count(&count).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	if count>0{
		return serializer.BuildChatlistResponse()
	}
	chatlist:=model.Chatlist{
		Sid:u.ID,
		Rid:service.Rid,
	}
	//创建
	if err:=model.DB.Create(&chatlist).Error;err!=nil{
		//添加日志
		logger.Log().Error("数据库创建用户失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildChatlistResponse()
}

