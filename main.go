package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine

func init()  {
	driveName := "mysql"
	DsName := "im:SYYGP2ZJfsssKaBB@(127.0.0.1:3306)/im?charset=utf8"
	DbEngin, e := xorm.NewEngine(driveName, DsName)
	if e != nil {
		log.Fatal(e.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//最大连接数
	DbEngin.SetMaxOpenConns(2)

	//自动建表 结构体User
	//DbEngin.Sync2(new(User))
}

func main() {

}