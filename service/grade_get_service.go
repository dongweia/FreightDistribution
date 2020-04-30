package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

//获取信用分（已完成订单数量）
type GradeGetService struct {
	Num uint `form:"num" json:"num" `  //0是客户/1为运货商
}
func (service *GradeGetService)Get(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言
	num:=0
	if service.Num==0{
		if err:=model.DB.Model(&model.Order{}).Where("uid=? and state=?",u.ID,model.Finish).Count(&num).Error;err!=nil{
			logger.Log().Error("数据库查询失败", err)
			return  serializer.DBErr("数据库操作失败",err)
		}
	}else if service.Num==1{
		if err:=model.DB.Table("orders").Select("orders.id").Joins("left join commodities on commodities.id = orders.cid").Where("commodities.mid=? and state=?",u.ID,model.Finish).Count(&num).Error;err!=nil{
			logger.Log().Error("数据库查询失败", err)
			return  serializer.DBErr("数据库操作失败",err)
		}
	}

	return serializer.BuildGradeResponse(num)
}