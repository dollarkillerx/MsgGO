/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:40
* @Description:
* */
package config

import (
	"MsgGO/auth/defs"
	"MsgGO/auth/execption"
	"encoding/json"
	"os"
)

const (
	configPath = "config.json"
)

var (
	ConfigBase *defs.Config
)

func init()  {
	file, e := os.Open(configPath)
	execption.ExecptionPanic(e)
	decoder := json.NewDecoder(file)

	ConfigBase = &defs.Config{}
	e = decoder.Decode(ConfigBase)
	execption.ExecptionPanic(e)
}

