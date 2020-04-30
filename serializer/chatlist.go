package serializer

import (
	"FreightDistribution/model"
)

// chatlist 聊天列表
type Chatlist struct {
	Id uint `json:"id"`
	CreatedAt int64 `json:"created_at"`
	LastContent string `json:"last_content"`
	Type uint  `json:"type"` //0普通消息 1图片
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
}

//BuildChatlist 序列化
func BuildChatlist(chatlist *[]model.GetChatlist)(*[]Chatlist)  {
	newchatlist:=make([]Chatlist,len(*chatlist))
	for i,v:=range *chatlist{
		newchatlist[i].Id=v.Id
		newchatlist[i].CreatedAt=v.CreatedAt.Unix()
		newchatlist[i].LastContent=v.LastContent
		newchatlist[i].Type=v.Type
		newchatlist[i].Nickname=v.Nickname
		newchatlist[i].Avatar=v.Avatar
	}
	return &newchatlist
}

//BuildChatlistResponse 序列化响应
func BuildChatlistResponse() *Response {
	return &Response{
		Data:"",
	}
}

//BuildChatlistResponse 序列化查询列表响应
func BuildALLChatlistResponse(list *[]model.GetChatlist) *Response {
	return &Response{
		Data:*BuildChatlist(list),
	}
}