/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:53
* @Description:
* */
package result

import (
	"MsgGO/auth/defs"
	"encoding/json"
	"net/http"
)

func ResultBase(w http.ResponseWriter,result *defs.BaseResult)  {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(result.HttpSc)
	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}