package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)

// TransporterGetService 查看商家信息服务
type TransporterGetService struct {
	ID uint `form:"id" json:"id"`
}


func (service *TransporterGetService)Get() *serializer.Response  {
	user:=model.User{}
	if err:=model.DB.First(&user, service.ID).Error;err!=nil{
		logger.Log().Error("数据库查询商家信息失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildTransporterResponse(&user)
}


