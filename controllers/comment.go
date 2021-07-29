package controllers

import (
	models "hello/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego_web "github.com/beego/beego/v2/server/web"
)

type CommentController struct {
	beego_web.Controller
}

func (c *CommentController) AddComment() {
	var comment models.Comment
	comment.GameId = c.GetString("GameId")
	comment.Content = c.GetString("Content")
	comment.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	// timeStr:=time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法

	o := orm.NewOrm()
	_, err := o.Insert(&comment)
	if err != nil {
		c.Ctx.Output.Body([]byte("err"))
	}
	redirectUrl := "/game/"
	redirectUrl += c.GetString("GameId")
	c.Ctx.Redirect(302, redirectUrl)
}
