package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"MsgGO/user/config"
	"MsgGO/user/exception"
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
}