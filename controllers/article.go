package controllers

import (
	"beego_study/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"os"
	"path"
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
	c.TplName = "article.tpl"
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
	// 限定文件格式
	fileExt := path.Ext(fh.Filename)
	fileExt = strings.ToLower(fileExt)
	if fileExt != ".jpg" && fileExt != ".png" {
		c.Data["message"] = "文件类型不符"
		return
	}
	// 限制文件大小
	if fh.Size > 20*1024*1024 {
		c.Data["message"] = "文件大小超过20M"
		return
	}
	// 定义新图片名字
	tn := time.Now().UnixNano()
	t := strconv.FormatInt(tn, 10)
	newName := t + fileExt
	fileDir, err := os.Getwd()
	logs.Info(fileDir)
	fileDir = fileDir + "/static/img"
	newPath := fileDir + "/" + newName
	err = c.SaveToFile("pic", newPath)
	if err != nil {
		c.Data["message"] = "移动图片失败"
		return
	}
	logs.Info(title, content, author, fh.Size, fh.Filename, fh.Header, newName)
	c.TplName = "article.tpl"
	if title == "" || content == "" || author == "" {
		c.Data["message"] = "有字段为空"
		return
	}
	o := orm.NewOrmUsingDB("test")
	a := &models.Article{}
	a.Title = title
	a.Content = content
	a.Author = author
	a.AddTime = time.Now()
	a.QiniuBucket = fileDir
	a.QiniuKey = newName
	id, err := o.Insert(a)
	logs.Info(id, err)
	if err != nil {
		c.Data["message"] = "插入数据库错误"
		return
	}
	c.Data["message"] = "写入文章成功，ID：" + strconv.FormatInt(id, 10)
}

// ArticleList 文章列表页
func (c *ArticleController) ArticleList() {
	c.TplName = "article_list.tpl"
	o := orm.NewOrmUsingDB("test")
	var articles []models.Article
	num, err := o.QueryTable(models.Article{}).All(&articles)
	if err != nil {
		logs.Info(err)
		return
	}
	logs.Info(num, articles)
	c.Data["list"] = articles
}
