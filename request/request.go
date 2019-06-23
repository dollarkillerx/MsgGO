/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午10:34
* */
package request

import (
	"MsgGO/defs"
	"encoding/json"
	"net/http"
)

func ReqMsg(w http.ResponseWriter, data *defs.RequestDate) {
	w.WriteHeader(data.Code)
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(data.Data)
	w.Write(bytes)
}

func ReqInpMsg(w http.ResponseWriter, code int, msg string) {
	date := &defs.RequestDate{Code: code, Data: &defs.SMap{"Message": msg}}
	ReqMsg(w, date)
}

func ReqInpData(w http.ResponseWriter, code int, data interface{}) {
	date := &defs.RequestDate{Code: code, Data: data}
	ReqMsg(w, date)
}
