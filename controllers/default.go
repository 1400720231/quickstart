package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

//beego的web Controller相当于django的View
type MainController struct {
	beego.Controller
}

//给beego的Controller动态添加一个方法get，这里的c相当于python的self，基础go语法
func (c *MainController) Get() {
	//这里之所以能用c.Data 是因为beego.Controller的结构体有这些属性方法
	//相当于python父类的属性
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//mvc 中的view 在Django中封装了这一层，就是html文件，给别人看的叫view
	c.TplName = "index.tpl"
}
