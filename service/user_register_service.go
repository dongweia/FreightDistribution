package service

import (
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/logger"
	"regexp"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Phone           string `form:"phone" json:"phone" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return serializer.ParamErr("两次输入的密码不相同",nil)
	}

	phoneOK,_:=regexp.MatchString(`^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$`,service.Phone)
	if !phoneOK {
		return serializer.ParamErr("输入的手机号格式不对",nil)
	}

	count := 0
	if err:=model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count).Error;err != nil {
		logger.Log().Error("数据库查询nickname失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	if count > 0 {
		return serializer.ParamErr("昵称被占用",nil)
	}

	count = 0
	if err:=model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count).Error;err != nil {
		logger.Log().Error("数据库查询user_name失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	if count > 0 {
		return serializer.ParamErr("用户名已经注册",nil)
	}

	count = 0
	if err:=model.DB.Model(&model.User{}).Where("phone = ?", service.Phone).Count(&count).Error;err != nil {
		logger.Log().Error("数据库查询phone失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	if count > 0 {
		return serializer.ParamErr("手机号已被注册",nil)
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
	// 表单验证
	if errresponse := service.valid(); errresponse != nil {
		return errresponse
	}

	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Phone: service.Phone,
		Status:   model.Active,
		Avatar:"https://dongwei-1300856266.cos.ap-chengdu.myqcloud.com/avatar/nickname.png",
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.EncryptErr()
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		//添加日志
		logger.Log().Error("数据库创建用户失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}

	return serializer.BuildUserResponse(&user)
}
