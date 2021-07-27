package controllers

import (
	"encoding/json"
	"fmt"
	models "hello/models"
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

func (g *GameListController) List() {
	parm_map := g.Ctx.Input.Params()
	var str string
	str = "game list list"
	fmt.Println(parm_map)
	json_parm, err := json.Marshal(parm_map)
	if err != nil {
		g.Ctx.Output.Body([]byte(err.Error()))
	}
	str = string(json_parm) + str
	g.Ctx.Output.Body([]byte(str))
}
