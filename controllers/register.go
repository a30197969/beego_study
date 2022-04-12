package controllers

import (
	"beego_study/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type RegisterController struct {
	beego.Controller
}

// 用户注册
func (c *RegisterController) Get() {
	c.TplName = "register.tpl"
}

// 提交注册
func (c *RegisterController) Post() {
	c.TplName = "register.tpl"
	// 获取POST数据
	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	if name == "" || password == "" || repassword == "" {
		c.Data["message"] = "用户名或者密码为空"
		return
	}
	if password != repassword {
		c.Data["message"] = "两次输入的密码不一致"
		return
	}
	id, err := models.InsertUser(name, password)
	if err != nil {
		c.Data["message"] = "写入数据库失败"
		logs.Info(err)
		return
	}
	c.Data["message"] = "注册成功，ID：" + strconv.FormatInt(id, 10)
	return
	// c.Redirect("/login", 302) // URL跳转
}

// 删除一条注册数据
func (c *RegisterController) RegisterDelete() {
	tmpId := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(tmpId, 10, 64)
	if err != nil {
		c.Ctx.WriteString("类型转换失败")
		return
	}
	o := orm.NewOrmUsingDB("test")
	u := &models.User{
		Id: id,
	}
	num, err := o.Delete(u)
	if err != nil {
		c.Ctx.WriteString("删除当前注册用户失败（" + tmpId + "）")
		return
	}
	c.Ctx.WriteString(strconv.FormatInt(num, 10))
}

// 更新一条注册数据
func (c *RegisterController) RegisterUpdate() {
	tmpId := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(tmpId, 10, 64)
	if err != nil {
		c.Ctx.WriteString("类型转换失败")
		return
	}
	o := orm.NewOrmUsingDB("test")
	u := &models.User{
		Id: id,
	}
	err = o.Read(u)
	if err != nil {
		c.Ctx.WriteString("当前注册用户不存在")
		return
	}
	dateline := strconv.FormatInt(time.Now().Unix(), 10)
	u.Name = "pengliang" + dateline
	u.Password = dateline
	num, err := o.Update(u, "name", "password")
	if err != nil {
		c.Ctx.WriteString("更改当前注册用户失败（" + tmpId + "）")
		return
	}
	c.Ctx.WriteString(strconv.FormatInt(num, 10))
	return
}

// 查询一条注册数据
func (c *RegisterController) RegisterOne() {
	c.EnableRender = false
	// 获取路由参数
	tmpId := c.Ctx.Input.Param(":id")
	logs.Info(tmpId)
	id, err := strconv.ParseInt(tmpId, 10, 64)
	if err != nil {
		c.Ctx.WriteString("类型转换失败")
		return
	}
	userInfo, err := models.SelectUserById(id)
	if err != nil {
		c.Ctx.WriteString("序列化失败")
		return
	}
	logs.Info(userInfo, err)
	var result string
	if userInfo != nil {
		resultByte, _ := json.Marshal(userInfo)
		result = string(resultByte)
	} else {
		result = "当前注册用户不存在"
	}
	c.Ctx.WriteString(result)
	return
}

// 查询所有注册用户
func (c *RegisterController) RegisterList() {
	c.EnableRender = false
	users, err := models.SelectUserList()
	if err != nil {
		c.Ctx.WriteString("查询所有注册用户失败")
		return
	}
	num := len(users)
	if num <= 0 {
		c.Ctx.WriteString("无数据")
		return
	}
	//for i := 0; i < len(users); i++ {
	//	result, _ := json.Marshal(users[i])
	//	c.Ctx.WriteString(string(result))
	//}
	result, _ := json.Marshal(users)
	c.Ctx.WriteString(string(result))
}
