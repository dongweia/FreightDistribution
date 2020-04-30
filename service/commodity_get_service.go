package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

// CommodityGetService 删除一个商品服务
type CommodityGetService struct {
	ID uint `form:"id" json:"id" `
}

func (service *CommodityGetService)Get(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return serializer.ParamErr("非法操作",nil)
	}
	u,_:=user.(*model.User) //类型断言

    comm:=model.Commodity{}
	if err :=model.DB.Where("id=? and mid=?",service.ID,u.ID).First(&comm).Error; err != nil {
		logger.Log().Error("数据库查询失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildGetCommodityResponse(&comm)
}

