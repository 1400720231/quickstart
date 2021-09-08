package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

//beego的web Controller相当于django的View
type MainController struct {
	beego.Controller
}

//这相当于django的某个视图函数，所以下面的Get()方法在UserController中只能出现一个，因为在http方法在一个请求中只能出现一次
//django中是cbv方式实现，会被覆盖不会报错，但是如果你有两个func (c *MainController) Get() {}这样的结构体go会报错重复声明
type UserController struct {
	beego.Controller
}

//给beego的Controller动态添加一个方法get，这里的c相当于python的self，基础go语法
func (c *MainController) Get() {
	//这里之所以能用c.Data 是因为beego.Controller的结构体有这些属性方法
	//相当于python父类的属性
	c.Data["Website"] = "beego.me" //相当于django的context=context
	c.Data["Email"] = "astaxie@gmail.com"
	//mvc 中的view 在Django中封装了这一层，就是html文件，给别人看的叫view
	c.TplName = "index.tpl"
}

//user路由对应的Controller，对应django中的viee或者tornado中的handler
//实现对应的get方法
func (c *UserController) Get() {
	//当前请求的url:/user
	fmt.Println(c.Ctx.Request.RequestURI)
	// user?name=xiongyao查询参数中的xiongyao
	fmt.Println(c.Ctx.Input.Query("name"))
	//获取查询参数 没有的话给默认值
	//在 /Users/xiongyao/go/pkg/mod/github.com/beego/beego/v2@v2.0.1/server/web/controller.go下面有
	//很多基于c.Ctx.Input.Query封装的获取默认值方法
	fmt.Println(c.GetString("name", "刘德华"))
	c.Data["code"] = 200 //相当于django的context=context
	c.Data["data"] = "panda"
	c.Data["message"] = "请求user成功"
	c.TplName = "user.tpl"

}
