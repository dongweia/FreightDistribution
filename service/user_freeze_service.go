package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"fmt"
)

// FreezeUserService 冻结服务
type FreezeUserService struct {
	Uid uint `form:"uid" json:"uid" `
}

//FreezeUser冻结
func (service *FreezeUserService)Freeze() *serializer.Response  {
	fmt.Println(service.Uid)
	if err:=model.DB.Model(&model.User{}).Where("id=?",service.Uid).Update("status",model.Suspend).Error;err!=nil{
		logger.Log().Error("数据库封禁账号失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildFreezeUserResponse()
}

