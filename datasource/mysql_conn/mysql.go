/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午11:21
* */
package mysql_conn

import (
	"MsgGO/config"
	"MsgGO/datamodels"
	"github.com/go-xorm/xorm"
)

var (
	MySQLEngine *xorm.Engine
	err         error
)

func init() {
	MySQLEngine, err = xorm.NewEngine("mysql", config.Config.MySQLDsn)
	if err != nil {
		panic(err.Error())
	}
	if config.Config.Debug {
		MySQLEngine.ShowSQL(true)
	}

	// 开启内存LRU缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	MySQLEngine.SetDefaultCacher(cacher)

	// 映射sql
	syncTab()
}

func syncTab() {
	err := MySQLEngine.Sync2(
		new(datamodels.User), // 用户表
	)
	if err != nil {
		panic(err.Error())
	}
}
