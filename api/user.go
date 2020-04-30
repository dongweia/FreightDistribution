package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"FreightDistribution/serializer"
	"FreightDistribution/service"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//UserPassword 用户修改密码
func UserPassword(c *gin.Context)  {
	var service service.UserPasswordService
	if err:=c.ShouldBind(&service);err==nil {
		res := service.UpdatePassword()
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}

}

//UserNickname 用户修改昵称
func UserNickname(c *gin.Context)  {
	var service service.UserUpdateNicknameService
	if err:=c.ShouldBind(&service);err==nil{
		res := service.UpdateNickname(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}



// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, *res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// UserAvatar 用户头像接口
func UserAvatar(c *gin.Context) {
	res := service.FileDeal(c)
	c.JSON(200, *res)
}

// FreezeUser 冻结用户接口
func FreezeUser(c *gin.Context) {
	var service service.FreezeUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Freeze()
		c.JSON(200, *res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}