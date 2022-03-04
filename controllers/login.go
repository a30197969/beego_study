package controllers

import (
	"beego_study/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.ViewPath = "views"
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	name := c.GetString("name")
	password := c.GetString("password")
	logs.Info(name, password)
	c.TplName = "login.tpl"
	if name == "" || password == "" {
		c.Data["message"] = "用户名或者密码为空"
		return
	}
	o := orm.NewOrmUsingDB("test")
	u := &models.User{}
	u.Name = name
	err := o.Read(u, "Name")
	logs.Info(u)
	if err != nil {
		c.Data["message"] = "当前用户不存在"
		return
	}
	if u.Password != password {
		c.Data["message"] = "登录密码错误"
		return
	}
	//c.Ctx.WriteString("登录成功")
	c.Redirect("/index", 302)
}
