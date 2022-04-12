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
	//uid := c.Ctx.GetCookie("beego_uid")
	uid := c.GetSession("beego_uid")
	if value, ok := uid.(int); ok {
		logs.Info(value, ok)
		if value > 0 {
			c.Redirect("/article_list", 302)
		}
	}
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
	// 设置cookie
	cokId := int(u.Id)
	//c.Ctx.SetCookie("beego_uid", cokId, time.Second*86400*3)
	c.SetSession("beego_uid", cokId)
	//c.Ctx.WriteString("登录成功")
	c.Redirect("/article_list", 302)
}
