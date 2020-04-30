package serializer

type Grade struct {
	Grade uint `json:"grade" `
}
// BuildGradeResponse 序列化分数响应
func BuildGradeResponse(g int) *Response {
	return &Response{
		Data:g,
	}
}