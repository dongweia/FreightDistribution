package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// addfriend
func Addfriend(c *gin.Context) {
	var service service.AddFriendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Add(c)
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// deletefriend
func Deletefriend(c *gin.Context) {
	var service service.DeleteFriendService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c)
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// addfriend
func Getallfriend(c *gin.Context) {
		res := service.GetAllFriend(c)
		c.JSON(200, *res)
}