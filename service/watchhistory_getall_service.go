package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"github.com/gin-gonic/gin"
)

func GetAllWatchHistory(c *gin.Context)*serializer.Response{
	list:=[]model.Commodity{}
	user,exist:=c.Get("user")
	if exist==false{
		return  serializer.NoRightErr("非法访问",nil)
	}
	u,_:=user.(*model.User) //类型断言select cid from collects where deleted_at IS NULL and mid=?
	//查询自己的发布商品
	if err:=model.DB.Raw("select * from commodities where id in (select cid from watch_histories where deleted_at IS NULL and uid=?) and deleted_at IS NULL",u.ID).Order("commodities.id desc").Scan(&list).Error;err!=nil{
		logger.Log().Error("数据库查询失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}

	return serializer.BuildGetAllWatchHistoryResponse(&list)
}

