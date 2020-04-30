package serializer

import "FreightDistribution/model"

// Commodity 商品序列化器
type Commodity struct {
	ID uint `json:"id"`
	Time int64 `json:"time"`
	Mid uint `json:"mid"`
	Title string `json:"title"`
	Describe string `json:"describe"`
	Image string `json:"image"`
	Lable string `json:"lable"`
	Address string `json:"address"`
}

// BuildCommoditylist 序列化商品列表
func BuildCommoditylist(commodity *[]model.Commodity) *[]Commodity {
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

// BuildCommodity 序列化商品信息
func BuildCommodity(commodity *model.Commodity) *Commodity {
	return &Commodity{
		ID:  commodity.ID,
		Time: commodity.UpdatedAt.Unix(),
		Mid:  commodity.Mid,
		Title:  commodity.Title,
		Describe: commodity.Describe,
		Image: commodity.Image,
		Lable: commodity.Lable,
		Address: commodity.Address,
	}
}


// BuildCommodityUploadResponse 序列化上传成功响应
func BuildCommodityUploadResponse() *Response {
	return &Response{
	}
}


// BuildCommodityDeleteResponse 序列化删除成功响应
func BuildCommodityDeleteResponse() *Response {
	return &Response{
	}
}

// BuildGetResponse 序列化获取成功响应
func BuildGetCommodityResponse(commodity *model.Commodity) *Response {
	return &Response{
		Data:BuildCommodity(commodity),
	}
}

// BuildGetAllResponse 序列化获取所有成功响应
func BuildGetAllCommodityResponse(commodity *[]model.Commodity) *Response {
	return &Response{
		Data:BuildCommoditylist(commodity),
	}
}
