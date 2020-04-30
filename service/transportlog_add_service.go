package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"

)

// TransportlogAddService 查看商家信息服务
type TransportlogAddService struct {
	Oid uint `form:"oid" json:"oid" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Phone string `form:"phone" json:"phone"`
}

func (service *TransportlogAddService)Add()*serializer.Response{
	transportlog:=model.Transportlog{
		Oid:service.Oid,
		Content:service.Content,
		Phone:service.Phone,
	}
	if err:=model.DB.Create(&transportlog).Error;err!=nil{
		//添加日志
		logger.Log().Error("数据库创建物流信息记录失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}

	return serializer.BuildTransportlogResponse()
}