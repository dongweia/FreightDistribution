package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//添加一个收藏
type AddCollectService struct {
	Cid uint `form:"cid" json:"cid" binding:"required"`
}

//添加
func (service *AddCollectService)Add(c *gin.Context) *serializer.Response  {
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	//验证是否重复
	count:=0
	if err:=model.DB.Model(&model.Collect{}).Where(&model.Collect{Uid:u.ID,Cid:service.Cid}).Count(&count).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	if count>0{
		return serializer.BuildAddCollectResponse(1) //1代表已存在
	}
	collect:=model.Collect{
		Uid:u.ID,
		Cid:service.Cid,
	}
	//创建
	if err:=model.DB.Create(&collect).Error;err!=nil{
		//添加日志
		logger.Log().Error("数据库创建收藏失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildAddCollectResponse(0)
}



