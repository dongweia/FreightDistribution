package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

// CommodityDeleteService 删除一个商品服务
type CommodityDeleteService struct {
	ID uint `form:"id" json:"id" `
}

func (service *CommodityDeleteService)Delete(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return serializer.ParamErr("非法操作",nil)
	}
	u,_:=user.(*model.User) //类型断言

	if err :=model.DB.Where("id = ? and mid = ?",service.ID,u.ID).Delete(&model.Commodity{}).Error; err != nil {
		logger.Log().Error("数据库删除失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildCommodityDeleteResponse()
}
