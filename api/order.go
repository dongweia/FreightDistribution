package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// Order 增添订单接口
func AddOrder(c *gin.Context) {
	var service service.AddOrderService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Add(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

//用户获取订单接口
func GetAllOrderUser(c *gin.Context) {
	var service service.GetAllOrderUserService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.GetAll(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

//商家获取订单接口
func GetAllOrderBusinesser(c *gin.Context) {
	var service service.GetAllOrderBusinessService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.GetAll(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

//商家获取订单接口
func UserOrderEnter(c *gin.Context) {
	var service service.UserOrderEnterService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Enter(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}