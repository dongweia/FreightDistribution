package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/uuid"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"os"
	"path"
	"time"
)

type File struct {
	Name string `form:"name" json:"name" binding:"required"`
	Size int64 `form:"size" json:"size" binding:"required,lte=10240"`   //单位kb  需小于10m
}
func(service *File) CosCredentialsService(c *gin.Context) *serializer.Response {
	user,exist:=c.Get("user")
	if exist==false{
		return serializer.ParamErr("非法操作",nil)
	}
	u,_:=user.(*model.User) //类型断言
	Suffix:=path.Ext(service.Name) //获取文件后缀
	newname:=uuid.Makeuuid()
	Credentials,err:=makeCredentials()
	if err!=nil {
		logger.Log().Error("生成cos秘钥失败", err)
		return serializer.CosErr()
	}
	cos:=model.Cos{
		Uid:u.ID,
		Name:newname+Suffix,
		Size:service.Size,
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
	return serializer.BuildCosResponse(&cos)
}

//调用腾讯cos接口生成证书
func makeCredentials() (*sts.Credentials,error){
	appid := os.Getenv("COS_APPID")
	bucket :=os.Getenv("COS_BUCKET")
	region :=os.Getenv("COS_REGION")
	c := sts.NewClient(
		os.Getenv("COS_SECRETID"),
		os.Getenv("COS_SECRETKEY"),
		nil,
	)
	//c := sts.NewClient(
	//	"AKIDl6A5COepZqqFDzi4HDn57yZyCB0kOUZH",
	//	"fneKuEeohJqKfd4eqKQkSu5yavCE81cp",
	//	nil,
	//)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          os.Getenv("COS_REGION"),
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/*",
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	if err != nil {
		return &sts.Credentials{},err
	}
	//fmt.Printf("%+v\n", res)
	//fmt.Printf("%+v\n", res.Credentials)
	return res.Credentials,nil
}