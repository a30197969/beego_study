package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "fengniao.com"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.ViewPath = "views"
}
func (c *MainController) Test() {
	c.EnableRender = false
	fmt.Println(c.Ctx.Request.RequestURI)
	fmt.Println(c.Ctx.Input.Query("name"))
	fmt.Println(c.Ctx.Input.Query("age"))
}
