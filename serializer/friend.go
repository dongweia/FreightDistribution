package serializer

import "FreightDistribution/model"

type FriendList struct {
	Id uint `json:"id" `
    Nickname string `json:"nickname" `
	Avatar string `json:"avatar" `
}

func BuildFriendList(friendlist *[]model.FriendDeal)*[]FriendList{
	newfriendlist:=make([]FriendList,len(*friendlist))
	for i,v:=range *friendlist{
		newfriendlist[i].Id=v.Id
		newfriendlist[i].Nickname=v.Nickname
		newfriendlist[i].Avatar=v.Avatar
	}
	return &newfriendlist
}
// BuildFriendResponse 序列化用户增加响应
func BuildFriendResponse(code int) *Response {
	return &Response{
		Code:code,
	}
}

////BuildFriendlistResponse 序列化查询列表响应
func BuildALLFriendlistResponse (list *[]model.FriendDeal)*Response{
	return &Response{
		Data:*BuildFriendList(list),
	}
}