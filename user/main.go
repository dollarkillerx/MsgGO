package main

import (
	"MsgGO/user/config"
	"MsgGO/user/exception"
	"MsgGO/user/router"
	"net/http"
)

func main() {
	// 注册路由
	router := router.RegisterRouter()

	err := http.ListenAndServe(config.BaseConfig.Host, router)
	exception.SimpleException(err)
}