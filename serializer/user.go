package serializer

import "FreightDistribution/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	Phone	string   `json:"phone"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
		Phone:	   hiddenphone(user.Phone), //隐藏中间四位
	}
}

//手机号隐藏中间4位
func hiddenphone(str string) string{
	ch:=[]byte(str)
	for i:=3;i<=6;i++{
		ch[i]='*'
	}
	return string(ch)
}


// BuildUserResponse 序列化用户响应
func BuildUserResponse(user *model.User) *Response {
	return &Response{
		Data: BuildUser(user),
	}
}

// BuildUserLoginSuccessResponse 序列化用户登陆成功响应
func BuildUserLoginSuccessResponse(user *model.User,token string) *Response {
	return &Response{
		Data: BuildUser(user),
		Token:token,
	}
}

// BuildFreezeUserResponse 序列化冻结用户用户响应
func BuildFreezeUserResponse() *Response {
	return &Response{
	}
}