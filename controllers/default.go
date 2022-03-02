package controllers

import (
	"beego_study/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
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

// Register 注册页面
func (c *MainController) Register() {
	//c.Data["name"] = "fengniao.com"
	c.TplName = "register.tpl"
}

// RegisterPost 提交注册
func (c *MainController) RegisterPost() {
	// 获取POST数据
	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	if name == "" || password == "" {
		logs.Warning("用户名或者密码为空")
		c.Redirect("/register", 302)
	}
	if password != repassword {
		logs.Warning("密码不一致")
		c.Redirect("/register", 302)
	}
	id, err := models.InsertUser(name, password)
	if err != nil {
		logs.Warning("插入数据库失败")
		c.Redirect("/register", 302)
	}
	c.Data["message"] = "注册成功，ID：" + strconv.FormatInt(id, 10)
	c.TplName = "register.tpl"

}

// 登录页面
func (c *MainController) Login() {
	//c.Data["name"] = "fengniao.com"
	c.TplName = "login.tpl"
}

func (c *MainController) LoginPost() {
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
	err := o.Read(u,"Name")
	fmt.Println(u)
	if err != nil {
		c.Data["message"] = "当前用户不存在"
		return
	}
	if u.Password != password {
		c.Data["message"] = "登录密码错误"
		return
	}
	c.Ctx.WriteString("登录成功")
}

// 删除一条数据
func (c *MainController) RegisterDel() {
	o := orm.NewOrmUsingDB("test")
	u := &models.User{
		Id: 10,
	}
	num, _ := o.Delete(u)
	c.Ctx.WriteString(strconv.FormatInt(num, 10))
}

// 更新一条数据
func (c *MainController) RegisterUpdate() {
	o := orm.NewOrmUsingDB("test")
	u := &models.User{
		Id: 6,
	}
	err := o.Read(u)
	if err != nil {
		c.Ctx.WriteString("当前数据不存在")
	}
	u.Name = "pengliang5"
	u.Password = "555555"
	num, _ := o.Update(u, "name", "password")
	c.Ctx.WriteString(strconv.FormatInt(num, 10))
}

// 查询一条注册数据
func (c *MainController) RegisterOne() {
	c.EnableRender = false
	tmpId := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(tmpId, 10, 64)
	user, err := models.SelectUser(id)
	if err != nil {
		c.Ctx.WriteString("序列化失败")
	}
	result, _ := json.Marshal(user)
	c.Ctx.WriteString(string(result))
}

// 查询所有注册用户
func (c *MainController) RegisterList() {
	c.EnableRender = false
	user, err := models.SelectUser(5)
	if err != nil {
		c.Ctx.WriteString("序列化失败")
	}
	result, _ := json.Marshal(user)
	c.Ctx.WriteString(string(result))
}

func (c *MainController) Test() {
	// 打印路由参数
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)

	c.Ctx.WriteString("hello")
	bbuserid := c.Ctx.GetCookie("bbuserid")
	fmt.Println(bbuserid)

	user, _ := beego.AppConfig.String("mysqluser")
	fmt.Println(user)

	c.EnableRender = false
	fmt.Println(c.GetString("hello"))
	fmt.Println(c.Ctx.Request.RequestURI)
	fmt.Println(c.Ctx.Input.Query("name"))
	fmt.Println(c.Ctx.Input.Query("age"))
	fmt.Println(c.Ctx.Input.Query("age"))
	cName, cAction := c.Controller.GetControllerAndAction()
	// 当你设置了自动渲染，然后在你的 Controller 中没有设置任何的 TplName，那么 beego 会自动设置你的模板文件如下：
	// 也就是你对应的 Controller 名字+请求方法名.模板后缀，也就是如果你的 Controller 名是 AddController
	// 请求方法是 POST，默认的文件后缀是 tpl，那么就会默认请求 /viewpath/addcontroller/post.tpl 文件
	tplName := strings.ToLower(cName) + "/" + strings.ToLower(cAction) + "." + c.TplExt
	fmt.Println(cName, cAction, tplName)
}
