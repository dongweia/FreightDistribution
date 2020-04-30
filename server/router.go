package server

import (
	"FreightDistribution/api"
	"FreightDistribution/middleware"
	"FreightDistribution/token"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	setMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	// 中间件, 顺序不能改
	r.Use(token.Verify)
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		//修改密码
		v1.PUT("user/password",api.UserPassword)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)  //个人中心
			auth.PUT("user/nickname",api.UserNickname) //修改昵称
			auth.POST("user/avatar",api.UserAvatar)

			auth.POST("user/chatlist",api.Addchat) //添加一个聊天
			auth.DELETE("user/chatlist",api.Deletechat)
			auth.GET("user/chatlist",api.GetAllchat)
			auth.GET("user/chathistory",api.GetChatHistory)
			auth.POST("user/chatimage",api.ChatImage)

			auth.POST("user/friend",api.Addfriend)//好友
			auth.GET("user/friendlist",api.Getallfriend)
			auth.DELETE("user/friend",api.Deletefriend)

			auth.GET("user/grade",api.Grade)//信用分

			auth.POST("user/opinion",api.OpinionAdd) //意见
			auth.GET("user/opinionlist",api.OpinionGetAll)

			auth.POST("commodity",api.AddCommodity) //商品
			auth.POST("updatecommodity",api.UpdateCommodity) //商品
			auth.DELETE("commodity",api.DeleteCommodity)
			auth.GET("commodity",api.GetCommodity)
			auth.GET("commoditylist",api.GetAllCommodity)

			auth.POST("commodity/order",api.AddOrder) //订单
			auth.GET("commodity/orderbusinesserlist",api.GetAllOrderBusinesser)
			auth.GET("commodity/orderuserlist",api.GetAllOrderUser)
			auth.PUT("commodity/order",api.UserOrderEnter)

			auth.POST("commodity/transportlog",api.Transportlog) //物流信息Transportloglist
			auth.GET("commodity/transportloglist",api.Transportloglist)

			auth.GET("transporter",api.Transporter)//商家信息

			auth.POST("commodity/collect",api.AddCollect) //收藏
			auth.GET("commodity/collectlist",api.GetAllCollect)
			auth.DELETE("commodity/collect",api.DeleteCollect)

			auth.POST("commodity/watchhistory",api.AddWatchHistory) //添加浏览记录
			auth.GET("commodity/watchhistorylist",api.GetAllWatchHistory)
			auth.DELETE("commodity/watchhistory",api.DeleteWatchHistory)





			//cos
			auth.GET("coscredentials",api.CosCredentials)

			//websocket
			auth.GET("user/websocket",api.WSchatlist)

			//管理员
			admin:= auth.Group("")
			admin.Use(middleware.Admin())
			{
				admin.GET("admin/me", api.UserMe)
				admin.GET("admin/chatlist",api.GetAllchat)
				admin.PUT("admin/freeze", api.FreezeUser)
				auth.DELETE("admin/commodity",api.AdminDeleteCommodity)

			}
		}
	}
	return r
}
