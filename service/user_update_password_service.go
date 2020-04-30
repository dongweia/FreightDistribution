package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)

// UserLoginService 管理用户登录的服务
type UserPasswordService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	OldPassword string `form:"old_password" json:"old_password" binding:"required,min=8,max=40"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

//UpdatePassword 修改密码
func (service *UserPasswordService)UpdatePassword() *serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return  serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.OldPassword) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if service.PasswordConfirm != service.Password {
		return serializer.ParamErr("新密码不相同",nil)
	}
	// 加密新密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.EncryptErr()
	}

	//更新到数据库
	if err:=model.DB.Save(&user).Error;err!=nil{
		logger.Log().Error("数据库更新失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildUserResponse(&user)

}



