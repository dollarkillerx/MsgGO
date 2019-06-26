/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-23
* Time: 上午10:56
* */
package main

import (
	"MsgGO/ws/container"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.GET("/ws",container.Ws)

	app.Run(":9001")
}
