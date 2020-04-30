package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// Collect 增添收藏接口
func AddCollect(c *gin.Context) {
	var service service.AddCollectService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Add(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

// Collect 删除收藏接口
func DeleteCollect(c *gin.Context) {
	var service service.CollectDeleteService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Delete(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取所有收藏接口
func GetAllCollect(c *gin.Context) {
	res := service.GetAllCollect(c)
	c.JSON(200, *res)
}