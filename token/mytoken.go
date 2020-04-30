package token

import (
	"FreightDistribution/logger"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	ErrorServerBusy = "server is busy"
	ErrorReLogin = "relogin"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

var (
	Secret = "123#111"  //salt
	ExpireTime = 7*24*60*60  //token expire time
)

//generate jwt token
func genToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorServerBusy)
	}
	return signedToken, nil
}

//登录，获取jwt token
func MakeToken(c *gin.Context,user_id uint)(string,error) {
	claims := &JWTClaims{
		UserID: user_id,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(claims)
	if err != nil {
		return "",err
	}
	return singedToken,nil

}

//验证jwt token
func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorServerBusy)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReLogin)
	}
	//fmt.Println("verify")
	return claims, nil
}

//验证
func Verify(c *gin.Context) {

	strToken := c.GetHeader("Authorization")

	if strToken!=""{
		claim, err := verifyAction(strToken)
		if err != nil {
			logger.Log().Error("token验证失败", err)
			return
		}
		c.Set("user_id",claim.UserID)
		return
	}
	strToken=c.DefaultQuery("token","")
	if strToken!=""{
		claim, err := verifyAction(strToken)
		if err != nil {
			logger.Log().Error("token验证失败", err)
			return
		}
		c.Set("user_id",claim.UserID)
		return
	}
}





//
//func sayHello(c *gin.Context) {
//	strToken := c.Param("token")
//	claim, err := verifyAction(strToken)
//	if err != nil {
//		c.String(http.StatusNotFound, err.Error())
//	}
//	c.String(http.StatusOK, "hello, ", claim.Username)
//}
//
//
//
//func refresh(c *gin.Context) {
//	strToken := c.Param("token")
//	claims, err := verifyAction(strToken)
//	if err != nil {
//		c.String(http.StatusNotFound, err.Error())
//		return
//	}
//	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
//	signedToken, err := genToken(claims)
//	if err != nil {
//		c.String(http.StatusNotFound, err.Error())
//		return
//	}
//	c.String(http.StatusOK, signedToken, ", ", claims.ExpiresAt)
//}