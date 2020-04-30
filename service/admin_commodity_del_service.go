package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)

// 	AdminCommodityDeleteService 删除一个商品服务
type AdminCommodityDeleteService struct {
	Cid uint `form:"cid" json:"cid" binding:"required" `
}

func (service *AdminCommodityDeleteService)Delete()*serializer.Response{
	if err :=model.DB.Where("id = ? ",service.Cid).Delete(&model.Commodity{}).Error; err != nil {
		logger.Log().Error("数据库删除失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildCommodityDeleteResponse()
}

