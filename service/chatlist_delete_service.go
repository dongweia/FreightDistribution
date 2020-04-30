package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//用户聊天列表 添加一个聊天
type DeleteChatService struct {
	Rid uint `form:"rid" json:"rid" binding:"required"`
}

//添加
func (service *DeleteChatService)Delete(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	//删除
	if err:=model.DB.Where("sid = ? and rid = ?",u.ID,service.Rid).Delete(&model.Chatlist{}).Error; err != nil {
		//添加日志
		logger.Log().Error("数据库删除用户失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildChatlistResponse()
}
