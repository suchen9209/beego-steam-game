package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) LoginPage() {
	c.TplName = "login.html"
}

func (c *UserController) Login() {

}
