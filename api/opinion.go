package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

//Opinion 意见添加
	func OpinionAdd(c *gin.Context)  {
	var service service.OpinionAddService
	if err:=c.ShouldBind(&service);err==nil{
		res := service.Add(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}


func OpinionGetAll(c *gin.Context)  {
		res := service.GetAllOpinion()
		c.JSON(200, *res)
}