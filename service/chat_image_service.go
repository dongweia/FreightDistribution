package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/uuid"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
)

func ChatImage(c *gin.Context)*serializer.Response {
	//得到上传的文件
	file, _, err := c.Request.FormFile("image") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		return serializer.ParamErr("接收文件失败",err)
	}

	//文件的名称
	filetxt :=path.Ext(c.Request.FormValue("name"))
	newname:=uuid.Makeuuid()+filetxt
	err=savephoto(&file,newname)
	if err!=nil{
		return serializer.ParamErr("创建文件失败",err)
	}

	size, err:=strconv.ParseInt(c.Request.FormValue("size"), 10, 64)

	user,exist:=c.Get("user")
	if exist==false{
		return serializer.ParamErr("非法操作",nil)
	}
	u,_:=user.(*model.User) //类型断言

	Credentials,err:=makeCredentials()
	if err!=nil {
		logger.Log().Error("生成cos秘钥失败", err)
		return serializer.CosErr()
	}
	cos:=model.Cos{
		Uid:u.ID,
		Name:newname,
		Size:size/1024,   //kb
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
	cospath,err:=uploadCOS(cos.TmpSecretID,cos.TmpSecretKey,cos.SessionToken,"chatimage",newname,"photo/"+newname)
	if err!=nil {
		return serializer.CosErr()
	}
	//可考虑移除 临时图片
	return serializer.BuildAvatarResponse(cospath)
}



