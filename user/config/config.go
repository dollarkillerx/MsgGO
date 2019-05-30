package config

import (
	"MsgGO/user/exception"
	"encoding/json"
	"os"
)

type baseCon struct {
	Host string `json:"host"` // 开放端口
	DriverName string `json:"driverName"` // db driver
	Dsn string `json:"dsn"` // dsn
	Test string `json:"test"`  // 测试开关
	PrivateKeyPath string `json:"private_key_path"` // rsa私钥path
	PublicKeyPath string `json:"public_key_path"` // rsa公钥path
}

const (
	filePath = "config.json"
)

var (
	BaseConfig *baseCon
)

// 初始化配置文件
func init()  {
	file, e := os.Open(filePath)
	exception.SimpleException(e)
	decoder := json.NewDecoder(file)

	BaseConfig = &baseCon{}
	e = decoder.Decode(BaseConfig)
	exception.SimpleException(e)
}

