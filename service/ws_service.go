package service

import (
	"FreightDistribution/logger"
	"FreightDistribution/model"
	"FreightDistribution/serializer"
	"FreightDistribution/wsconn"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

//以服务器时间为准
var upGrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin:func(r *http.Request) bool{
		return true
	},
}
//发送来的数据
type WBSData struct {
	Rid uint `json:"rid" form:"rid"`
	Content string `json:"content" form:"content"`
	Type uint  `json:"type" form:"type"`
}


func WSservice(c *gin.Context)  {
	user,exist:=c.Get("user")
	if exist==false{
		return
	}
	u,_:=user.(*model.User) //类型断言
	// 完成ws协议的握手操作
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	fmt.Println("用户ID：",u.ID,"已连接")

	wsconn.SM.Store(u.ID,ws)

	defer fmt.Println("用户ID：",u.ID,"连接退出")
	defer wsconn.SM.Delete(u.ID)
	defer ws.Close()
	for  {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//接收json
		var wbsdata WBSData
		if err := json.Unmarshal([]byte(message), &wbsdata); err != nil{
			break
		}
        fmt.Println(mt)
		//对发来的消息进行处理
		//rid 为0是测试连接是否断开 否者进行处理
		if wbsdata.Rid!=0{
			wbsdata.Deal(u.ID)
		}


        returnmess:=serializer.ReturnSendOK()
		//写入ws数据
		err = ws.WriteMessage(mt, returnmess)
		if err != nil {
			break
		}
	}
}

func (wb *WBSData)Deal(sid uint)  {
	v,ok:=wsconn.SM.Load(wb.Rid)
	if ok==false{//说明对方不在线
		//直接入库

		tx:=model.DB.Begin()
		//删除前面 插入两条新数据
		if err := tx.Where("sid = ? and rid = ?",sid,wb.Rid).Or("sid = ? and rid = ?",wb.Rid,sid).Delete(&model.Chatlist{}).Error; err != nil {
			logger.Log().Error("数据库操作失败1", err)
			tx.Rollback()
			return
		}
		if err := tx.Create(&model.Chatlist{Sid:sid,Rid:wb.Rid,LastContent:wb.Content,Type:wb.Type}).Error; err != nil {
			logger.Log().Error("数据库操作失败2", err)
			tx.Rollback()
			return
		}
		if err := tx.Create(&model.Chatlist{Sid:wb.Rid,Rid:sid,LastContent:wb.Content,Type:wb.Type}).Error; err != nil {
			logger.Log().Error("数据库操作失败3", err)
			tx.Rollback()
			return
		}
		if err := tx.Create(&model.Chatmess{Sid:sid, Rid:wb.Rid, Content:wb.Content, Status:0, Type:wb.Type,}).Error; err != nil {
			logger.Log().Error("数据库操作失败4", err)
			tx.Rollback()
			return
		}
		tx.Commit()

	}else {//在线情况
      conn,_:=v.(*websocket.Conn)
      var ruser model.User
      if err := model.DB.First(&ruser,sid).Error;err!=nil{
			logger.Log().Error("数据库操作失败0", err)
			return
      }
      fmt.Println(ruser)
      mess:=(&serializer.WBRdata{
		   Code:200,
		   Sid:ruser.ID,
		   Sname:ruser.Nickname,
		   Savatar:ruser.Avatar,
		   Content:wb.Content,
		   Type:wb.Type,
		   Time:time.Now().Unix(),
	   }).SendMess()

		//写入ws数据
		err := conn.WriteMessage(1, mess)
		if err != nil {
			return
		}
		tx:=model.DB.Begin()
		//删除前面 插入两条新数据
		if err := tx.Where("sid = ? and rid = ?",sid,wb.Rid).Or("sid = ? and rid = ?",wb.Rid,sid).Delete(&model.Chatlist{}).Error; err != nil {
			logger.Log().Error("数据库操作失败1", err)
			tx.Rollback()
			return
		}
		if err := tx.Create(&model.Chatlist{Sid:sid,Rid:wb.Rid,LastContent:wb.Content,Type:wb.Type}).Error; err != nil {
			logger.Log().Error("数据库操作失败2", err)
			tx.Rollback()
			return
		}
		if err := tx.Create(&model.Chatlist{Sid:wb.Rid,Rid:sid,LastContent:wb.Content,Type:wb.Type}).Error; err != nil {
			logger.Log().Error("数据库操作失败3", err)
			tx.Rollback()
			return
		}
		if err := tx.Create(&model.Chatmess{Sid:sid, Rid:wb.Rid, Content:wb.Content, Status:1, Type:wb.Type,}).Error; err != nil {
			logger.Log().Error("数据库操作失败4", err)
			tx.Rollback()
			return
		}
		tx.Commit()

	}
}