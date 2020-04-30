package serializer

import "FreightDistribution/model"

//BuildAddWatchHistoryResponse 序列化增加一个商品浏览响应
func BuildAddWatchHistoryResponse() *Response {
	return &Response{
	}
}

//BuildGetAllWatchHistoryResponse 序列化获取所有浏览响应
func BuildGetAllWatchHistoryResponse(commodity *[]model.Commodity) *Response {
	return &Response{
		Data:BuildCollect(commodity),
	}
}

//BuildDeleteWatchHistoryResponse 序列化删除一个浏览响应
func BuildDeleteWatchHistoryResponse() *Response {
	return &Response{
	}
}