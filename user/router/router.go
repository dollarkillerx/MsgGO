package router

import (
	"MsgGO/user/container"
	"github.com/julienschmidt/httprouter"
)

func RegisterRouter() *httprouter.Router {

	router := httprouter.New()

	router.POST("/user/userlogin",container.UserLogin) // 用户登陆
	router.POST("/user/registered",container.UserRegistered) // 用户注册

	return router
}
