package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.TplName = "article.tpl"
}

func (c *ArticleController) Post() {
	title := c.GetString("title")
	content := c.GetString("content")
	author := c.GetString("author")
	// 获取图片
	f, fh, err := c.GetFile("pic")
	defer f.Close()
	if err != nil {
		c.Data["message"] = "上传文件失败"
		return
	}
	tn := time.Now().UnixNano()
	t := strconv.FormatInt(tn, 10)
	pos := strings.LastIndex(fh.Filename, ".")
	newName := t + fh.Filename[pos:]
	c.SaveToFile("pic", "/tmp/"+newName)
	logs.Info(title, content, author, fh.Size, fh.Filename, fh.Header, err, newName)
	c.TplName = "article.tpl"
	if title == "" || content == "" || author == "" {
		c.Data["message"] = "有字段为空"
		return
	}

	//o := orm.NewOrmUsingDB("test")
	//u := &models.User{}
	//u.Name = name
	//err := o.Read(u, "Name")
	//logs.Info(u)
	//if err != nil {
	//	c.Data["message"] = "当前用户不存在"
	//	return
	//}
	//if u.Password != password {
	//	c.Data["message"] = "登录密码错误"
	//	return
	//}
	////c.Ctx.WriteString("登录成功")
	//c.Redirect("/index", 302)
}
