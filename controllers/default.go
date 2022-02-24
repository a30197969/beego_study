package controllers

import (
	"beego_study/models"
	"fmt"
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
	c.Data["name"] = "fengniao.com"
	c.TplName = "register.tpl"
}

// RegisterPost 提交注册
func (c *MainController) RegisterPost() {
	// 获取POST数据
	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	if password != repassword {
		c.Ctx.WriteString("密码不一致")
	}
	id, err := models.InsertUser(name, password)
	if err != nil {
		c.Ctx.WriteString("注册用户失败")
	}
	c.Data["message"] = "注册成功，ID：" + strconv.FormatInt(id, 10)
	c.TplName = "register.tpl"

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
