package serializer

import "FreightDistribution/model"

// Collect 收藏商品序列化器
type Collect struct {
	ID uint `json:"id"`
	Time int64 `json:"time"`
	Mid uint `json:"mid"`
	Title string `json:"title"`
	Describe string `json:"describe"`
	Image string `json:"image"`
	Lable string `json:"lable"`
	Address string `json:"address"`
}
// BuildCollect 序列化收藏商品列表
func BuildCollect(commodity *[]model.Commodity) *[]Commodity {
	newcommoditylist:=make([]Commodity,len(*commodity))
	for i,v:=range *commodity{
		newcommoditylist[i].ID=v.ID
		newcommoditylist[i].Time=v.UpdatedAt.Unix()
		newcommoditylist[i].Mid=v.Mid
		newcommoditylist[i].Title=v.Title
		newcommoditylist[i].Describe=v.Describe
		newcommoditylist[i].Image=v.Image
		newcommoditylist[i].Lable=v.Lable
		newcommoditylist[i].Address=v.Address
	}
	return &newcommoditylist
}

//BuildGetAllCollectResponse 序列化获取所有收藏响应
func BuildGetAllCollectResponse(commodity *[]model.Commodity) *Response {
	return &Response{
		Data:BuildCollect(commodity),
	}
}

//BuildAddCollectResponse 序列化增加一个收藏响应
func BuildAddCollectResponse(code int) *Response {
	return &Response{
		Code:code,
	}
}

//BuildDeleteCollectResponse 序列化删除一个收藏响应
func BuildDeleteCollectResponse() *Response {
	return &Response{
	}
}