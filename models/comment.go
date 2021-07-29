package models

import "github.com/beego/beego/v2/client/orm"

type Comment struct {
	Id         int
	Content    string
	GameId     string
	CreateTime string
	WriterId   string
}

func init() {
	// set default database
	// orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:3306)/steam_game?charset=utf8&loc=Local")

	// register model
	// orm.RegisterModel(new(Game))

	// create table
	// orm.RunSyncdb("default", false, true)
	orm.RegisterModel(new(Comment))
}
