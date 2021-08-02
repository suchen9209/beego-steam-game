package controllers

import (
	models "hello/models"
	"math"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type GameController struct {
	beego.Controller
}

func (g *GameController) AddGame() {
	g.TplName = "add_game.html"
}

func (g *GameController) Get() {
	id := g.Ctx.Input.Param(":id")
	idint, _ := strconv.Atoi(id)
	game, err := models.GetGameInfo(idint)
	if err != nil {
		g.Abort("404")
	}
	comments := models.GetCommentList(idint)

	g.Data["game"] = game
	g.Data["Comments"] = comments
	g.TplName = "game_detail.html"
}

func (g *GameController) Post() {
	var game models.Game
	g.ParseForm(&game)
	game_id, _ := models.AddGame(&game)
	g.Ctx.Redirect(302, "/game/"+strconv.Itoa(int(game_id)))
}

type GameListController struct {
	beego.Controller
}

func (g *GameListController) Get() {
	g.Ctx.Output.Body([]byte("game list"))
}

type ListGame struct {
	Id       int
	GameName string
	Desc     string
}

func (g *GameListController) List() {
	parm_map := g.Ctx.Input.Params()
	per_page := 10
	page := parm_map["0"]
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		panic("Get out here!You monster!")
	}
	count := models.GetCount()

	// qb, _ := orm.NewQueryBuilder("mysql")
	// qb.Select("id", "game_name", "`desc`").
	// 	From("game").
	// 	OrderBy("id").Desc().
	// 	Limit(per_page).Offset((pageInt - 1) * per_page)
	// sql := qb.String()
	// fmt.Println(sql)
	// var games []models.Game
	// var games []ListGame
	// res := orm.Params{}
	// dataNum, _ := o.Raw("select id, game_name from game order").RowsToMap(&res, "id", "game_name")
	// qqr := o.QueryTable("game").OrderBy("id").Limit(per_page).Offset((pageInt - 1) * per_page)
	// dataNum, _ := qqr.All(&games)
	// for key, value := range res {
	// 	tmp := new(models.Game)
	// 	tmp.Id, _ = strconv.Atoi(key)
	// 	tmp.GameName = value.(string)
	// 	games = append(games, *tmp)
	// }

	// dataNum, _ := o.Raw(sql).QueryRows(&games)
	games, dataNum := models.GetGameList(pageInt, per_page)
	page_number := int(math.Ceil(float64(count) / float64(per_page)))

	// fmt.Println(res)

	g.Data["AllCount"] = count
	g.Data["AllPage"] = page_number
	g.Data["NowPage"] = pageInt
	g.Data["DataNumber"] = dataNum
	g.Data["List"] = games
	if pageInt-1 <= 0 {
		g.Data["Prev"] = 1
	} else {
		g.Data["Prev"] = pageInt - 1
	}
	if pageInt+1 >= page_number {
		g.Data["Next"] = page_number
	} else {
		g.Data["Next"] = pageInt + 1
	}

	g.TplName = "game_list.html"

}
