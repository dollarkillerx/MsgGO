/**
* @Author: dollarkiller
* @Date: 19-5-29 上午9:35
* @Description:
* */
package test

import (
	"MsgGO/user/utils"
	"encoding/base64"
	"io/ioutil"
	"os/exec"
	"runtime"
	"testing"
)

func TestOs(t *testing.T)  {
	s := runtime.GOOS
	t.Log(s)
}

func TestEx(t *testing.T) {
	command := exec.Command("openssl", "genrsa -out app_private_key.pem   1024")
	e1 := command.Start()
	if e1 != nil{
		t.Fatal(e1.Error())
	}
	cmd := exec.Command("openssl", "rsa -in app_private_key.pem -pubout -out app_public_key.pem")
	e2 := cmd.Start()
	if e2 != nil{
		t.Fatal(e2.Error())
	}
}

func TestRSA(t *testing.T) {
	bytes, e := ioutil.ReadFile("../app_public_key.pem")
	if e != nil {
		t.Fatal(e.Error())
	}
	ordata := []byte("123456")
	encrypt, e := utils.Rsa256Encrypt(ordata, bytes)
	if e != nil {
		t.Fatal(e.Error())
	}
	s := base64.URLEncoding.EncodeToString(encrypt)

	file, e := ioutil.ReadFile("../app_private_key.pem")
	if e != nil {
		t.Fatal(e.Error())
	}

	decodeString, _ := base64.URLEncoding.DecodeString(s)
	decrypt, e := utils.Rsa256Decrypt([]byte(decodeString), file)
	if e != nil {
		t.Fatal(e.Error())
	}


	t.Log(s)
	t.Log(encrypt)
	t.Log(string(ordata))
	t.Log(string(decrypt))
}
