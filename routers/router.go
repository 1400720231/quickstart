package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"quickstart/controllers"
)

func init() {
	//路由
	beego.Router("/", &controllers.MainController{})
}
