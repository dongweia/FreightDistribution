package middleware

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid ,_:= c.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err != nil {
				logger.Log().Error("数据库查找用户失败", err)
			}else {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if u, ok := user.(*model.User); ok{
				if u.Status!=model.Suspend{
					c.Next()
					return
				}else {
					c.JSON(200, serializer.CheckFreeze())
					c.Abort()
				}

			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort() //中间件里面有错误如果不想继续后续接口的调用不能直接return，而是应该调用c.Abort()方法
	}
}

//管理员认证
func Admin()gin.HandlerFunc  {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if u, ok := user.(*model.User); ok {
				count:=0
				if err:=model.DB.Model(&model.Admin{}).Where("uid = ?",u.ID).Count(&count).Error;err != nil {
					logger.Log().Error("数据库查询uid失败", err)
					c.JSON(200,serializer.DBErr("数据库操作失败",err))
					c.Abort()
				}
				if count>0{
					c.Next()
					return
				}

			}
		}
		c.JSON(200, serializer.CheckAdmin())
		c.Abort() //中间件里面有错误如果不想继续后续接口的调用不能直接return，而是应该调用c.Abort()方法
	}
}