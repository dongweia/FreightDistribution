package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// Transportlog 物流详情
func Transportlog(c *gin.Context) {
	var service service.TransportlogAddService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Add()
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// Transportlog 物流详情列表
func Transportloglist(c *gin.Context) {
	var service service.TransportlogGetAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetAll()
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
