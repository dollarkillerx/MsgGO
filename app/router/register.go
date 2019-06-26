/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:06
* */
package router

import (
	"MsgGO/app/container"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	engine := gin.New()

	engine.POST("/register/app",container.RegisterApp)

	return engine
}
