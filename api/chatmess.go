package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// getchathistory
func GetChatHistory(c *gin.Context) {
	var service service.GetChatHistoryService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get(c)
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ChatImage 聊天图片接口
func ChatImage(c *gin.Context) {
	res := service.ChatImage(c)
	c.JSON(200, *res)
}