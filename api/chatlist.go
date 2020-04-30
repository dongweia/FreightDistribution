package api

import (
	"FreightDistribution/service"
	"github.com/gin-gonic/gin"
)

// addchat
func Addchat(c *gin.Context) {
	var service service.AddChatService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Add(c)
		c.JSON(200, *res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
//deletechat
func Deletechat(c *gin.Context)  {
	var service service.DeleteChatService
	if err:=c.ShouldBind(&service);err==nil{
		res:=service.Delete(c)
		c.JSON(200,*res)
	}else {
		c.JSON(200,ErrorResponse(err))
	}

}
//getallchat
func GetAllchat(c *gin.Context)  {
	res := service.GetAllChat(c)
	c.JSON(200, *res)
}

//websocket服务
func WSchatlist(c *gin.Context)  {
	service.WSservice(c)
}