/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:06
* */
package main

import (
	"MsgGO/app/router"
	"MsgGO/common/config"
)

func main() {
	app := router.RegisterRouter()

	app.Run(config.AppConfig.AppConf.Host)
}

