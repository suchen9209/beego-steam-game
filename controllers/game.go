package controllers

import (
	"fmt"
	models "hello/models"
	"math"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
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
	idint, err := strconv.Atoi(id)
	fmt.Println(idint)
	if err != nil {
		panic("id must be integer")
	}
	game := models.Game{
		Id: idint,
	}
	o := orm.NewOrm()
	o.Read(&game)

	g.Data["Id"] = game.Id
	g.Data["Name"] = game.GameName
	g.Data["Link"] = game.Link
	g.Data["Desc"] = game.Desc
	g.Data["Platform"] = game.Platform
	g.Data["GamePlat"] = game.Gameplat
	g.TplName = "game_detail.html"
	// gamename := g.Ctx.Input.Param(":gamename")

	// g.Ctx.Output.Body(s)
}

func (g *GameController) Post() {
	var game models.Game
	game.GameName = g.GetString("gameName")
	game.Link = g.GetString("gameLink")
	game.Desc = g.GetString("gameDesc")
	game.Platform = g.GetString("platform")
	game.Gameplat = g.GetString("gamePlat")

	o := orm.NewOrm()
	game_id, err := o.Insert(&game)
	if err != nil {
		g.Ctx.Output.Body([]byte("err"))
	}
	redirectUrl := "/game/"
	redirectUrl += strconv.Itoa(int(game_id))
	g.Ctx.Redirect(302, redirectUrl)

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
	o := orm.NewOrm()
	qr := o.QueryTable("game")
	count, _ := qr.Count()

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "game_name", "`desc`").
		From("game").
		OrderBy("id").Desc().
		Limit(per_page).Offset((pageInt - 1) * per_page)
	sql := qb.String()
	fmt.Println(sql)
	var games []models.Game
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

	dataNum, _ := o.Raw(sql).QueryRows(&games)
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
