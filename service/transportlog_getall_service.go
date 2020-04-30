package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)
// TransportlogGetAllService 查看物流信息服务
type TransportlogGetAllService struct {
	Oid uint `form:"oid" json:"oid" binding:"required"`
}
func (service *TransportlogGetAllService) GetAll()*serializer.Response{
	list:=[]model.Transportlog{}
	if err:=model.DB.Model(&model.Transportlog{}).Where("oid=?",service.Oid).Order("created_at desc").Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildTransportloglistResponse(&list)
}