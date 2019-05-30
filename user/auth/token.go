/**
* @Author: dollarkiller
* @Date: 19-5-29 上午9:18
* @Description:
* */
package auth

import (
	"MsgGO/user/config"
	"MsgGO/user/dbops/defsModel"
	"MsgGO/user/defs"
	"MsgGO/user/exception"
	"MsgGO/user/utils"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"strconv"
)

// 自动生成RSA公钥 私钥
func Rsa() {
// 在linux上安装好openssh不然会出错阿！
	command := exec.Command("openssl", "genrsa -out app_private_key.pem   1024")
	e1 := command.Start()
	if e1 != nil{
		panic(e1.Error())
	}
	cmd := exec.Command("openssl", "rsa -in app_private_key.pem -pubout -out app_public_key.pem")
	e2 := cmd.Start()
	if e2 != nil{
		panic(e2.Error())

}}

func GenerateToken(user *defsModel.ImUser) (string,error) {
	time := utils.GetCurrentTime()
	exps := 8*60*60
	i, err := strconv.Atoi(time)
	if err != nil{
		return "",err
	}
	exp := strconv.Itoa( i + exps)
	payloadi := &defs.Payload{User: user.Name, Iat: time,Exp:exp}
	headeri := &defs.Header{Type: "jwt"}

	payloadib, _ := json.Marshal(payloadi)
	headerb, _ := json.Marshal(headeri)

	payload := base64.URLEncoding.EncodeToString(payloadib)
	header := base64.URLEncoding.EncodeToString(headerb)

	hrs := header + "." + payload
	//privateKey, rerr := ioutil.ReadFile(config.BaseConfig.PrivateKeyPath)
	publicKey, uerr := ioutil.ReadFile(config.BaseConfig.PublicKeyPath)
	exception.SimpleException(uerr)

	bytes, err := utils.Rsa256Encrypt([]byte(hrs), publicKey)
	if err != nil {
		return "",err
	}
	s := hrs + "." + base64.URLEncoding.EncodeToString(bytes)
	return s,nil
}
