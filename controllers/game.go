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

	s, err := json.Marshal(&game)
	if err != nil {
		g.Ctx.Output.Body([]byte("err"))
	}

	// gamename := g.Ctx.Input.Param(":gamename")

	g.Ctx.Output.Body(s)
}

func (g *GameController) Post() {
	id := g.Ctx.Input.Param(":id")
	idint, err := strconv.Atoi(id)
	fmt.Println(idint)
	if err != nil {
		panic("id must be integer")
	}
	game := models.GetGameInfo(idint)

	s, err := json.Marshal(&game)
	if err != nil {
		g.Ctx.Output.Body([]byte("err"))
	}

	g.Ctx.Output.Body(s)
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
