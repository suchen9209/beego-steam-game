package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Game struct {
	Id       int
	GameName string
	Link     string
	Desc     string
	Platform string
	Gameplat string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:3306)/steam_game?charset=utf8&loc=Local")

	// register model
	orm.RegisterModel(new(Game))

	// create table
	// orm.RunSyncdb("default", false, true)
}

func GetGameInfo(id int) *Game {
	o := orm.NewOrm()
	game := Game{Id: id}
	err := o.Read(&game)
	if err != nil {
		panic("no id")
	}
	return &game

}

// func main() {
// 	orm.Debug = true
// 	// var w io.Writer
// 	// orm.DebugLog = orm.NewLog(w)

// 	// orm.RegisterModel(new(Game))
// 	// db1, err := orm.GetDB("default")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	panic("wtf")
// 	// }
// 	// fmt.Println(db1)

// 	// row := db1.QueryRow("select * from game")
// 	// fmt.Println(&row)

// 	o := orm.NewOrm()
// 	var game Game
// 	// game := new(Game)
// 	game.GameName = "dota2"
// 	game.Link = "https://www.steam.com/dota2"
// 	game.Gameplat = "steam"
// 	game.Platform = "PC"
// 	id, err := o.Insert(&game)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(id)

// }
