package serializer

// BuildUserResponse 序列化用户响应
func BuildAvatarResponse(path string) *Response {
	return &Response{
		Data:path,
	}
}
