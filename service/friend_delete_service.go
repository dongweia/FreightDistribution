package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//用户聊天列表 添加一个聊天
type DeleteFriendService struct {
	Oid uint `form:"oid" json:"oid" binding:"required"`
}

//添加
func (service *DeleteFriendService)Delete(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	//验证是否重复
	if err:=model.DB.Where("mid = ? and oid = ?",u.ID,service.Oid).Delete(&model.Friend{}).Error; err != nil {
		//添加日志
		logger.Log().Error("数据库删除用户失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildFriendResponse(0)
}


