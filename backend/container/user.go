/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午10:17
* */
package container

import (
	"MsgGO/defs"
	"MsgGO/request"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	if username == "" || password == "" {
		request.ReqMsg(w, defs.ErrorBadRequest)
		return
	}
	if username == "123" && password == "123" {
		request.ReqInpData(w, 200, defs.SMap{
			"token": "Token....",
		})
		return
	}
	request.ReqMsg(w, defs.ErrorBadRequest)
}

func UserLoginShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
