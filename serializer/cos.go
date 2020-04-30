package serializer

import "FreightDistribution/model"

//Cos 序列化
type Cos struct {
	Name string `json:"name"`
	TmpSecretID string `json:"tmp_secret_id"`
	TmpSecretKey string `json:"tmp_secret_key"`
	SessionToken string `json:"session_token"`
}


func BuildCos(cos *model.Cos) Cos {
	return Cos{
		Name:cos.Name,
		TmpSecretID:cos.TmpSecretID,
		TmpSecretKey:cos.TmpSecretKey,
		SessionToken:cos.SessionToken,
	}
}

//序列化响应
func BuildCosResponse(cos *model.Cos) *Response {
	return &Response{
		Data: BuildCos(cos),
	}
}
