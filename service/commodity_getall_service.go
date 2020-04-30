package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
	"strings"
)

// CommodityGetAllService 获得所有商品服务
type CommodityGetAllService struct {
	Mid uint `form:"mid" json:"mid" `
	Startaddress string  `form:"startaddress" json:"startaddress" `
	Endaddress  string  `form:"endaddress" json:"endaddress" `
}

func (service *CommodityGetAllService)GetAll(c *gin.Context)*serializer.Response{
	list:=[]model.Commodity{}
	if service.Mid==0 &&service.Startaddress!="" &&service.Endaddress!=""{
		sa:=strings.Split(service.Startaddress, ",")
		ea:=strings.Split(service.Endaddress, ",")
		if err:=model.DB.Raw(
			" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and FIND_IN_SET(?,address) and deleted_at IS NULL union " +
				" SELECT * FROM `commodities` WHERE FIND_IN_SET(?,address) and deleted_at IS NULL ",
				sa[0]+sa[1]+sa[2],ea[0]+ea[1]+ea[2],
				sa[0]+sa[1]+sa[2],ea[0]+ea[1],
				sa[0]+sa[1],ea[0]+ea[1]+ea[2],
				sa[0]+sa[1],ea[0]+ea[1],
				sa[0]+sa[1]+sa[2],ea[0],
				sa[0],ea[0]+ea[1]+ea[2],
				sa[0],ea[0]+ea[1],
				sa[0]+sa[1],ea[0],
				sa[0],ea[0],
				"全国",
			).Scan(&list).Error;err!=nil{
			logger.Log().Error("数据库查询失败", err)
			return  serializer.DBErr("数据库操作失败",err)
		}
	}else if service.Mid==0 && service.Startaddress== "" && service.Endaddress== "" {
		user,exist:=c.Get("user")
		if exist==false{
			return  serializer.NoRightErr("非法访问",nil)
		}
		u,_:=user.(*model.User) //类型断言
		//查询自己的发布商品
		if err:=model.DB.Model(model.Commodity{}).Where("mid=?",u.ID).Scan(&list).Error;err!=nil{
			logger.Log().Error("数据库查询失败", err)
			return  serializer.DBErr("数据库操作失败",err)
		}
	}else {
		//查询某人的发布商品
		if err:=model.DB.Model(model.Commodity{}).Where("mid=?",service.Mid).Scan(&list).Error;err!=nil{
			logger.Log().Error("数据库查询失败", err)
			return  serializer.DBErr("数据库操作失败",err)
		}
	}
	return serializer.BuildGetAllCommodityResponse(&list)
}