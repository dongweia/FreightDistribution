package serializer

import (
	"encoding/json"
)

//接收端收到的mess
type WBRdata struct {
	Code int `json:"code" form:"code"`
	Sid uint `json:"sid" form:"sid"`
	Sname string `json:"sname" form:"sname"`
	Savatar string `json:"savatar" form:"savatar"`
	Content string `json:"content" form:"content"`
	Type uint  `json:"type" form:"type"`
	Time int64 `json:"time" form:"time"`
}

//发送成功 返回消息
type SendReturn struct {
	Code int `json:"code" form:"code"`
	Mess string `json:"mess" form:"mess"`
}

func ReturnSendOK() []byte {
	jsonStr,_ :=json.Marshal(SendReturn{
		Code: 0,
		Mess: "successful",
	})
	return jsonStr
}

func (wb *WBRdata)SendMess()[]byte{
	jsonStr,_ :=json.Marshal(*wb)
	return jsonStr
}