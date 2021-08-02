package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	engine, _ := beego.AppConfig.String("defaultDataBase")
	mysqluser, _ := beego.AppConfig.String("mysqluser")
	mysqlurls, _ := beego.AppConfig.String("mysqlurls")
	steam_game, _ := beego.AppConfig.String("mysqldb")
	// set default database
	err := orm.RegisterDataBase("default", engine, mysqluser+"@tcp("+mysqlurls+":3306)/"+steam_game+"?charset=utf8&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	// register model
	// orm.RegisterModel(new(Game))

	// create table
	// orm.RunSyncdb("default", false, true)
	orm.RegisterModel(new(Comment))
	orm.RegisterModel(new(Game))
	o = orm.NewOrm()
}
