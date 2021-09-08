package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"quickstart/controllers"
)

func init() {
	//路由
	beego.Router("/", &controllers.MainController{})
	// user路由
	beego.Router("/user", &controllers.UserController{})
	// login路由
	beego.Router("/login.tpl", &controllers.LoginController{})
}
