package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/uuid"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

// CommodityAddService 上传一个商品服务
type CommodityAddService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=50"`
	Describe string  `form:"describe" json:"describe" binding:"required,max=1000"`
	Image  multipart.FileHeader  `form:"image" json:"image" `  //文件
	Lable string  `form:"lable" json:"lable" `
	Address string  `form:"address" json:"address" binding:"max=2000"`
}

func (service *CommodityAddService)Add(c *gin.Context)*serializer.Response{
	user,exist:=c.Get("user")
	if exist==false{
		return serializer.ParamErr("非法操作",nil)
	}
	u,_:=user.(*model.User) //类型断言

	//得到上传的文件
	file, header, err := c.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		return serializer.ParamErr("接收文件失败",err)
	}



	//文件的名称
	_, fileext := filepath.Split(header.Header.Get("Content-Type")) //目前还没想到好方法截取
	newname:=uuid.Makeuuid()+"."+fileext
	err=savephoto(&file,newname)
	if err!=nil{
		return serializer.ParamErr("创建文件失败",err)
	}

	//得到上传的数据
	Credentials,err:=makeCredentials()
	if err!=nil {
		logger.Log().Error("生成cos秘钥失败", err)
		return serializer.CosErr()
	}
	cos:=model.Cos{
		Uid:u.ID,
		Name:newname,
		Size:header.Size/1024,   //kb
		TmpSecretID:Credentials.TmpSecretID,
		TmpSecretKey:Credentials.TmpSecretKey,
		SessionToken:Credentials.SessionToken,
	}
	// 创建cos
	if err := model.DB.Create(&cos).Error; err != nil {
		//添加日志
		logger.Log().Error("数据库创建cos失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	cospath,err:=uploadCOS(cos.TmpSecretID,cos.TmpSecretKey,cos.SessionToken,"commodity",newname,"photo/"+newname)
	if err!=nil {
		return serializer.CosErr()
	}

	comm:=&model.Commodity{
		Mid:u.ID,
		Title:service.Title,
		Describe:service.Describe,
		Image:cospath,
		Lable:service.Lable,
		Address:service.Address,
	}
	// 创建commodity
	if err := model.DB.Create(&comm).Error; err != nil {
		//添加日志
		logger.Log().Error("数据库创建commodity失败", err)
		return serializer.DBErr("数据库操作失败",err)
	}
	return serializer.BuildCommodityUploadResponse()
}