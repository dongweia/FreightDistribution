package serializer

import "FreightDistribution/model"

// Transporter 商家信息序列化器
type Transporter struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	Phone	string   `json:"phone"`
}

// BuildUser 序列化用户
func BuildTransporter(user *model.User) Transporter {
	return Transporter{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
		Phone:	  user.Phone, //隐藏中间四位
	}
}

// BuildUserResponse 序列化用户响应
func BuildTransporterResponse(user *model.User) *Response {
	return &Response{
		Data: BuildTransporter(user),
	}
}
