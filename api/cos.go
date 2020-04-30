package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

func CosCredentials(c *gin.Context)  {
	var service service.File
	if err := c.ShouldBind(&service); err == nil {
		res := service.CosCredentialsService(c)
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}