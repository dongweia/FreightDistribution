package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// Transporter 商家详情
func Transporter(c *gin.Context) {
	var service service.TransporterGetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, *res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}