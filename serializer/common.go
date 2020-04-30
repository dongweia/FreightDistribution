package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"` //omitempty序列化json时若是0值忽略
	Token string      `json:"token,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义

// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	CodeCheckAdmin=40002
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}
// CheckFreeze 检查是否冻结
func CheckFreeze() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "已冻结",
	}
}

// CheckAdmin 检查管理员
func CheckAdmin() Response {
	return Response{
		Code: CodeCheckAdmin,
		Msg:  "非管理员",
	}
}
// Err 通用错误处理
func Err(errCode int, msg string, err error) *Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode { //设置gin模式 修改模式
		res.Error = err.Error()
	}
	return &res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) *Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

//CodeNoRightErr 未授权访问
func NoRightErr(msg string, err error)  *Response{
	if msg == "" {
		msg = "非法访问"
	}
	return Err(CodeNoRightErr, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) *Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}

// EncryptErr 加密失败
func EncryptErr() *Response {
	msg := "密码加密失败"
	return Err(CodeParamErr, msg, nil)
}

//生成cos证书失败
func CosErr() *Response {
	msg:="生成cos访问秘钥失败"
	return Err(CodeParamErr, msg, nil)
}