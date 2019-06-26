/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-25
* Time: 上午11:09
* */
package mysql_conn

import (
	"MsgGO/common/config"
	"github.com/go-xorm/xorm"
)

var (
	MySQL *xorm.Engine
)

func GetMySQL() *xorm.Engine {
	if MySQL != nil {
		return MySQL
	}
	engine, e := xorm.NewEngine("mysql", config.AppConfig.MySQLConf.Dsn)
	if e != nil {
		panic(e.Error())
	}
	MySQL = engine
	return MySQL
}



