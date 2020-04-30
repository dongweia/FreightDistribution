package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//添加一个意见反馈
type OpinionAddService struct {
	Content string `form:"content" json:"content" binding:"required,min=10,max=1000"`
}

//添加
func (service *OpinionAddService)Add(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	//创建
	if err:=model.DB.Create(&model.Opinion{Uid:u.ID, Content:service.Content,}).Error;err!=nil{
		//添加日志
		logger.Log().Error("数据库创建意见失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildOpinionAddResponse()
}



