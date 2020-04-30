package service

import (
	"github.com/gin-gonic/gin"
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/token"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) *serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return  serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	if user.Status==model.Suspend {
		return serializer.ParamErr("账号已冻结", nil)
	}

	token,err:=token.MakeToken(c,user.ID)
	if err!=nil{
		logger.Log().Error("token生成失败", err)
		return serializer.BuildUserResponse(&user)
	}
	return serializer.BuildUserLoginSuccessResponse(&user,token)
}
