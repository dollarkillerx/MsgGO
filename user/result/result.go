package result

import (
	"MsgGO/user/defs"
	"MsgGO/user/exception"
	"encoding/json"
	"net/http"
)

func RequestError(w http.ResponseWriter,err defs.ErrorResponse) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(err.HttpSC)
	bytes, e := json.Marshal(err.Error)
	exception.SimpleExceptionLog(e)
	_, e = w.Write(bytes)
	exception.SimpleExceptionLog(e)
}

func RequestSuccess(w http.ResponseWriter,success defs.SuccessResponse) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(success.HttpSC)
	bytes, e := json.Marshal(success.Data)
	exception.SimpleExceptionLog(e)
	_, e = w.Write(bytes)
	exception.SimpleExceptionLog(e)
}

func RequestInterFace(w http.ResponseWriter,sc int,data *map[string]interface{}) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(sc)
	bytes, e := json.Marshal(data)
	exception.SimpleExceptionLog(e)
	_, e = w.Write(bytes)
	exception.SimpleExceptionLog(e)
}