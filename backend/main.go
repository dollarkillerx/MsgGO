/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午10:11
* */
package main

import (
	"MsgGO/backend/router"
	"fmt"
	"net/http"
)

func main() {
	app := router.RegisterRouter()

	fmt.Println("app run http://127.0.0.1:8085 ")
	if err := http.ListenAndServe(":8085", app); err != nil {
		panic(err.Error())
	}
}
