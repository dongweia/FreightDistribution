package service

import (
	"github.com/gin-gonic/gin"
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)

// UserUpdateNickname 用户修改昵称服务
type UserUpdateNicknameService struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
}

//UpdateAvatar  修改头像
func (service *UserUpdateNicknameService)UpdateNickname(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
		if exist==false{
			return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	if err:=model.DB.Model(u).Update("nickname",service.Nickname).Error;err!=nil{
		logger.Log().Error("数据库更新昵称失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildUserResponse(u)
}

