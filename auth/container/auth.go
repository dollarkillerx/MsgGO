/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:37
* @Description:
* */
package container

import (
	"MsgGO/auth/config"
	"MsgGO/auth/defs"
	"MsgGO/auth/execption"
	"MsgGO/auth/result"
	"MsgGO/auth/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strings"
)

func AuthToken(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	r.ParseForm()
	token := r.PostForm.Get("token")
	if token == "" {
		result.ResultBase(w,defs.CheckAuthError)
	}
	split := strings.Split(token, ".")
	data := split[1]
	ciphertext := split[2]
	s, e := decrypt(ciphertext, data)
	if e != nil {
		result.ResultBase(w,defs.CheckAuthError)
	}
	switch s {
	case defs.TokenInvalid :
		result.ResultBase(w,defs.CheckAuthBeOverdue)
	case defs.TokenOk:
		result.ResultBase(w,defs.CheckAuthOK)
	}
}

func decrypt(str string,data string) (string,error) {
	bytes, e := ioutil.ReadFile(config.ConfigBase.PrivateKeyPath)
	execption.ExecptionPanic(e)
	decodeString, e := base64.URLEncoding.DecodeString(str)
	if e != nil {
		return defs.TokenOverdue,e
	}
	_, e = utils.Rsa256Decrypt([]byte(decodeString), bytes)
	if e != nil {
		return defs.TokenInvalid,e
	}

	i, _ := base64.URLEncoding.DecodeString(data)
	payload := &defs.Payload{}
	e = json.Unmarshal(i, payload)
	if e != nil {
		return "",e
	}
	time := utils.GetCurrentTime()
	fmt.Println(time)
	fmt.Println(payload.Exp)
	if time > payload.Exp {
		return defs.TokenInvalid,nil
	}else{
		return defs.TokenOk,nil
	}
}
