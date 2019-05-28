package config

import (
	"MsgGO/user/exception"
	"encoding/json"
	"os"
)

type baseCon struct {
	Host string `json:"host"`
	DriverName string `json:"driverName"`
	Dsn string `json:"dsn"`
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

