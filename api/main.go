package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}
// ErrorResponse 返回错误消息
func ErrorResponse(err error) *serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("提交参数不合法", err)
}