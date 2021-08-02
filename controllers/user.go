package controllers

import (
	"hello/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) LoginPage() {
	c.TplName = "login.html"
}

func (c *UserController) Login() {
	var user models.User
	c.ParseForm(&user)
	if models.CheckUser(&user) {
		c.SetSession("user_id", user.Id)
		c.Ctx.Redirect(302, "/game/new")
	} else {
		c.Abort("403")
	}
}
