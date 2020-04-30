package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
)

func GetAllOpinion()*serializer.Response{
	list:=[]model.OpinionUser{}
	//查询自己的发布商品
	if err:=model.DB.Table("opinions").Select("users.id as uid,users.nickname as uname,users.avatar as uavatar,opinions.content as content,opinions.created_at as created_at").Joins("left join users on users.id = opinions.uid").Order("opinions.created_at desc").Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildOpinionGetAllResponse(&list)
}
