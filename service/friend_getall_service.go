package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

func GetAllFriend(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	list:=[]model.FriendDeal{}
	if err:=model.DB.Table("friends").Select("friends.oid as id,users.nickname as nickname,users.avatar as avatar ").Joins("left join users on friends.oid = users.id").Where("friends.mid=?",u.ID).Order("friends.created_at").Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildALLFriendlistResponse(&list)
}