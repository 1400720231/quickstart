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
type LoginController struct {
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
	//当前请求的url:     /user
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

type LoginForm struct {
	/*
		//注意这里需要转换格式 表示把大写的Username专为username
		后面的`form:"username"`表示struct的tag标签用法，理论上可以随便取，但是我们这用的ParseForm是专门针对
		表单请求的解析，源码里面值针对form tag做了解析，如果你这里设置成`json:"username"`  ParseForm是解析不了的
		那么你就拿不到前端穿过来的post参数
	*/
	Username string `form:"username"`
	Password string `form:"password"`
}

func (c *LoginController) Post() {
	//当前请求的url:     /login.tpl
	fmt.Println(c.Ctx.Request.RequestURI)
	// user?name=xiongyao查询参数中的xiongyao
	fmt.Println(c.Ctx.Input.Query("name"))
	//这样也能取到post中的值 因为它的顺序是如果查询参数input.Params，取不到就会是form中取
	//username := c.GetString("username")
	//password := c.GetString("password")
	//
	//fmt.Println(username)
	//fmt.Println(password)
	//这样取这正取post中的值
	u := LoginForm{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
		fmt.Println("xxxx")
	}
	fmt.Println(u)
	//当然你要是想要在上面一层的form 请使用c.Ctx.Request.PostForm，在网上一层就是http协议中的body了
	//因为http层面是没有form概念的只有body

	c.Data["code"] = 200 //相当于django的context=context
	c.Data["username"] = u.Username
	c.Data["password"] = u.Password
	c.TplName = "login.tpl"

}
