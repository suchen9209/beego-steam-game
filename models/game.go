package models

type Game struct {
	Id       int    `form:"-"`
	GameName string `form:"gameName"`
	Link     string `form:"gameLink"`
	Desc     string `form:"gameDesc"`
	Platform string `form:"platform"`
	Gameplat string `form:"gamePlat"`
}

const table_name = "game"

// func init() {
// 	// set default database
// 	// orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:3306)/steam_game?charset=utf8&loc=Local")

// 	// register model

// 	// create table
// 	// orm.RunSyncdb("default", false, true)
// }

func GetGameInfo(id int) (*Game, error) {
	game := Game{Id: id}
	err := o.Read(&game)
	return &game, err
}

func AddGame(game *Game) (int64, error) {
	return o.Insert(&game)
}

func GetCount() int {
	qr := o.QueryTable(table_name)
	count, _ := qr.Count()
	return int(count)
}

func GetGameList(page int, per_page int) (*[]Game, int) {
	var games []Game
	qqr := o.QueryTable("game").OrderBy("id").Limit(per_page).Offset((page - 1) * per_page)
	dataNum, _ := qqr.All(&games)
	return &games, int(dataNum)
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
