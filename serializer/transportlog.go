package serializer

import "FreightDistribution/model"

// Transportlog 物流信息序列化器
type Transportlog struct {
	Oid        uint   `json:"oid"`
	Time  int64 `json:"time"`
	Content string `json:"content"`
	Phone string `json:"phone"`
}

// BuildTransportlog 序列化物流信息列表
func BuildTransportlog(transportlog *[]model.Transportlog) *[]Transportlog {
	newtransportlog:=make([]Transportlog,len(*transportlog))
	for i,v:=range *transportlog{
		newtransportlog[i].Oid=v.Oid
		newtransportlog[i].Time=v.CreatedAt.Unix()
		newtransportlog[i].Content=v.Content
		newtransportlog[i].Phone=v.Phone
	}
	return &newtransportlog
}

// BuildTransportloglistResponse 序列化添加物流信息响应
func BuildTransportloglistResponse(transportlog *[]model.Transportlog) *Response {
	return &Response{
		Data:BuildTransportlog(transportlog),
	}
}

// BuildTransportlogResponse 序列化添加物流信息响应
func BuildTransportlogResponse() *Response {
	return &Response{
	}
}
