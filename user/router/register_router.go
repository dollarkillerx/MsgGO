/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-23
* Time: 上午10:47
* */
package router

import "github.com/gin-gonic/gin"

func RegisterRouter() *gin.Engine {
	app := gin.New()


	AppRouter(app)

	return app
}