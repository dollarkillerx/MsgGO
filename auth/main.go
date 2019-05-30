/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:34
* @Description:
* */
package main

import (
	"MsgGO/auth/config"
	"MsgGO/auth/execption"
	"MsgGO/auth/router"
	"net/http"
)

func main() {
	router := router.RegisterRouter()

	err := http.ListenAndServe(config.ConfigBase.Host, router)
	execption.ExecptionPanic(err)
}
