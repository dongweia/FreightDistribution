package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func FileDeal(c *gin.Context)*serializer.Response {
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
	cospath,err:=uploadCOS(cos.TmpSecretID,cos.TmpSecretKey,cos.SessionToken,"avatar",newname,"photo/"+newname)
	if err!=nil {
		return serializer.CosErr()
	}
	if err:=model.DB.Model(u).Update("avatar",cospath).Error;err!=nil{
		logger.Log().Error("数据库更新头像失败", err)
		return  serializer.DBErr("数据库操作失败",err)
	}
	//可考虑移除 临时图片
	return serializer.BuildAvatarResponse(cospath)
}

func savephoto(file *multipart.File,name string)error{
	out, err := os.Create("photo/"+name)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, *file)
	if err != nil {
		return err
	}
	return nil
}

func uploadCOS(id string,key string,token string,class string,name string,file string)(string,error){
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	u, _ := url.Parse("https://dongwei-1300856266.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	//c := cos.NewClient(b, &http.Client{
	//	Transport: &cos.AuthorizationTransport{
	//		SecretID:  "COS_SECRETID",
	//		SecretKey: "COS_SECRETKEY",
	//	},
	//})

	// 2.临时密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:id,
			SecretKey:key,
			SessionToken:token,
		},
	})
	if client != nil {
		// 调用 COS 请求
	}
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//name := "avatar/01.jpg"
	// 1.通过字符串上传对象
	//f := strings.NewReader("test")
	//
	//_, err := client.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	panic(err)
	//}
	// 2.通过本地文件上传对象
	_, err := client.Object.PutFromFile(context.Background(),class+"/"+name, file, nil)
	if err != nil {
		return "",err
	}
	return "https://dongwei-1300856266.cos.ap-chengdu.myqcloud.com/"+class+"/"+name,nil
}
