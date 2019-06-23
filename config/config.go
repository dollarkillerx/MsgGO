/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午11:21
* */
package config

import (
	"encoding/json"
	"os"
)

type Conf struct {
	Host     string `json:"host"`
	MySQLDsn string `json:"mysql_dsn"`
	Debug    bool   `json:"debug"`
}

var (
	Config *Conf
)

func init() {
	path := "./config.json"
	file, e := os.Open(path)
	if e != nil {
		panic(e.Error())
	}

	Config = &Conf{}
	decoder := json.NewDecoder(file)
	e = decoder.Decode(Config)
	if e != nil {
		panic(e.Error())
	}
}
