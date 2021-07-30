package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:3306)/steam_game?charset=utf8&loc=Local")
	o = orm.NewOrm()
	// register model
	// orm.RegisterModel(new(Game))

	// create table
	// orm.RunSyncdb("default", false, true)
	orm.RegisterModel(new(Comment))
	orm.RegisterModel(new(Game))

}
