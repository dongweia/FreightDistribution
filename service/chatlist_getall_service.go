package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)


func GetAllChat(c *gin.Context) *serializer.Response  {
	//list:=[]model.Chatlist{}
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	list:=[]model.GetChatlist{}
	if err:=model.DB.Table("chatlists").Select("chatlists.rid as id,chatlists.created_at as created_at,chatlists.last_content as last_content,chatlists.type as type,users.nickname as nickname,users.avatar as avatar ").Joins("left join users on chatlists.rid = users.id").Where("chatlists.sid=?",u.ID).Where("chatlists.sid !=chatlists.rid").Order("chatlists.created_at desc").Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}

	//if err:=model.DB.Model(&model.Chatlist{}).Where(&model.Chatlist{Sid:u.ID,}).Order("created_at desc").Find(&list).Error;err!=nil{
	//	logger.Log().Error("数据库查询失败", err)
	//	return  serializer.DBErr("数据库操作失败",err)
	//}
	return serializer.BuildALLChatlistResponse(&list)
}