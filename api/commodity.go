package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// AddCommodity 增添商品接口
func AddCommodity(c *gin.Context) {
	var service service.CommodityAddService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Add(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteCommodity 删除商品接口
func DeleteCommodity(c *gin.Context) {
	var service service.CommodityDeleteService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Delete(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}
// AdminDeleteCommodity 删除商品接口
func AdminDeleteCommodity(c *gin.Context) {
	var service service.AdminCommodityDeleteService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Delete()
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetCommodity 获取商品接口
func GetCommodity(c *gin.Context) {
	var service service.CommodityGetService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Get(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}


// GetAllCommodity 获取所有商品接口
func GetAllCommodity(c *gin.Context) {
	var service service.CommodityGetAllService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.GetAll(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateCommodity 更新商品接口
func UpdateCommodity(c *gin.Context) {
	var service service.CommodityUpdateService
	if err:=c.ShouldBind(&service);err == nil {
		res := service.Update(c)
		c.JSON(200, *res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}
}