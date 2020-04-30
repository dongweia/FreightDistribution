package serializer

import "FreightDistribution/model"

type ChatHistory struct {
	Sid uint `json:"sid" form:"sid"`
	Sname string `json:"sname" form:"sname"`
	Savatar string `json:"savatar" form:"savatar"`
	Content string `json:"content" form:"content"`
	Type uint  `json:"type" form:"type"`
	Time int64 `json:"time" form:"time"`
}

//BuildChatlist 序列化
func BuildChatHistory(chathistory *[]model.ChatHistory)([]ChatHistory)  {
	newchathistory:=make([]ChatHistory,len(*chathistory))
	for i,v:=range *chathistory{
		newchathistory[i].Sid=v.Sid
		newchathistory[i].Sname=v.Sname
		newchathistory[i].Savatar=v.Savatar
		newchathistory[i].Content=v.Content
		newchathistory[i].Type=v.Type
		newchathistory[i].Time=v.Time.Unix()
	}
	return newchathistory
}

//BuildChatHistoryResponse 序列化响应
func BuildChatHistorytResponse(chathistory *[]model.ChatHistory) *Response {
	return &Response{
		Data:BuildChatHistory(chathistory),
	}
}
