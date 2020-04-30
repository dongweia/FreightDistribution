package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// AddWatchHistory 增添浏览记录接口
func AddWatchHistory(c *gin.Context) {
	var service service.AddWatchHistoryService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Add(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteWatchHistory  删除浏览接口
func DeleteWatchHistory(c *gin.Context) {
	var service service.WatchHistoryDeleteService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Delete(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取所有浏览接口
func GetAllWatchHistory(c *gin.Context) {
	res := service.GetAllWatchHistory(c)
	c.JSON(200, *res)
}