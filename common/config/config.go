/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:12
* */
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type App struct {
	Host string `yaml:"host"`
}

type MySQL struct {
	Dsn string `yaml:"dsn"`
}

type Config struct {
	AppConf App `yaml:"app"`
	MySQLConf MySQL `yaml:"mysql"`
}

var (
	AppConfig *Config
)

func init() {
	AppConfig = &Config{}
	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, AppConfig)
	if e != nil {
		panic(e.Error())
	}
}
