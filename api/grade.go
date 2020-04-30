package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// Grade 获取信用分接口
func Grade(c *gin.Context) {
	var service service.GradeGetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get(c)
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

