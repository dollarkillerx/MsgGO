package dbops

import (
	"MsgGO/user/config"
	"MsgGO/user/dbops/defsModel"
	"MsgGO/user/exception"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	err error
)

// 初始化数据库
func init() {
	Engine, err = xorm.NewEngine(config.BaseConfig.DriverName, config.BaseConfig.Dsn)
	exception.SimpleException(err)

	// 测试开关
	if config.BaseConfig.Test == "true" {
		Engine.ShowSQL(true)
	}

	syncDb()
}

// 数据库映射
func syncDb() {
	// 用户表映射
	err := Engine.Sync2(new(defsModel.ImUser))
	exception.SimpleException(err)

}
