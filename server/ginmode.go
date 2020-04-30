package server

import "github.com/gin-gonic/gin"

func setMode(model string) {

	initModel := gin.ReleaseMode //参数可以传递：gin.DebugMode、gin.ReleaseMode、gin.TestMode。
	switch model {
	case "release":
		initModel = gin.ReleaseMode
	case "debug":
		initModel = gin.DebugMode
	case "test":
		initModel = gin.TestMode
	}
	gin.SetMode(initModel)
}
