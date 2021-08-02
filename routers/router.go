package routers

import (
	"hello/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/game/:id([0-9]+)", &controllers.GameController{})
	beego.Router("/game/new", &controllers.GameController{}, "get:AddGame")
	// beego.Router("/game/:gamename:string", &controllers.GameController{})
	beego.Router("/game", &controllers.GameController{})
	beego.Router("/game_list", &controllers.GameListController{})

	beego.Router("/comment/new", &controllers.CommentController{}, "post:AddComment")
	beego.Router("/login", &controllers.UserController{}, "get:LoginPage")
	beego.Router("/login", &controllers.UserController{}, "post:Login")

	// beego.Router("/gaga", &controllers.MainController{})
	// beego.Router("/api/food",&RestController{},"get:ListFood;post:CreateFood;put:UpdateFood;delete:DeleteFood")

	beego.AutoRouter(&controllers.GameListController{})
}
