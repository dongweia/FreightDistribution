package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//查看聊天历史
type GetChatHistoryService struct {
	Id  uint `form:"id" json:"id" binding:"required"`
	Limit  uint `form:"limit" json:"limit" `
	Offset uint `form:"offset" json:"offset" `
}


func (service *GetChatHistoryService)Get(c *gin.Context) *serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	list:=[]model.ChatHistory{}
	if err:=model.DB.Raw("Select chatmesses.sid as sid,chatmesses.content as content,chatmesses.type as type ,chatmesses.created_at as time,users.nickname as sname,users.avatar as savatar from chatmesses left join users on chatmesses.sid = users.id where chatmesses.sid=? and chatmesses.rid=? union Select chatmesses.sid as sid,chatmesses.content as content,chatmesses.type as type ,chatmesses.created_at as time,users.nickname as sname,users.avatar as savatar from chatmesses left join users on chatmesses.sid = users.id where chatmesses.sid=? and chatmesses.rid=?  ORDER BY time desc limit ? offset ?",u.ID,service.Id,service.Id,u.ID,service.Limit,service.Offset).Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildChatHistorytResponse(&list)
}