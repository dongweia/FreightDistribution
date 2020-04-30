package serializer

import "FreightDistribution/model"

type OpinionUser struct {
	Uid uint `json:"uid"`
	Uname string `json:"uname"`
	Uavatar string `json:"uavatar"`
	Content string `json:"content"`
	Time int64 `json:"time"`
}


// BuildOpinionAddResponse 序列化删除成功响应
func BuildOpinionAddResponse() *Response {
	return &Response{
	}
}

// BuildOpinionUser 序列化反馈列表
func BuildOpinionUser(opinionuser *[]model.OpinionUser) *[]OpinionUser {
	newopinionuser:=make([]OpinionUser,len(*opinionuser))
	for i,v:=range *opinionuser{
		newopinionuser[i].Uid=v.Uid
		newopinionuser[i].Uname=v.Uname
		newopinionuser[i].Uavatar=v.Uavatar
		newopinionuser[i].Content=v.Content
		newopinionuser[i].Time=v.CreatedAt.Unix()
	}
	return &newopinionuser
}

// BuildOpinionGetAllResponse 序列化删除成功响应
func BuildOpinionGetAllResponse(opinionuser *[]model.OpinionUser) *Response {
	return &Response{
		Data:BuildOpinionUser(opinionuser),
	}
}