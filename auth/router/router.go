/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:36
* @Description:
* */
package router

import (
	"MsgGO/auth/container"
	"github.com/julienschmidt/httprouter"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/auth",container.AuthToken)

	return router
}
