package result

import (
	"MsgGO/user/defs"
	"MsgGO/user/exception"
	"encoding/json"
	"net/http"
)

func RequestError(w http.ResponseWriter,err defs.ErrorResponse) {
	w.WriteHeader(err.HttpSC)
	bytes, e := json.Marshal(err)
	exception.SimpleExceptionLog(e)
	_, e = w.Write(bytes)
	exception.SimpleExceptionLog(e)
}
