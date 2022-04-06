package controllers

import (
	"beego_study/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"math"
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
	o := orm.NewOrmUsingDB("test")
	var articleTypes []models.ArticleType
	o.QueryTable(models.ArticleType{}).All(&articleTypes)
	c.Data["articleTypes"] = articleTypes
	c.Data["headTitle"] = "添加文章页"
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "html_head.tpl"
	c.TplName = "article.tpl"
}

func (c *ArticleController) Post() {
	c.TplName = "article.tpl"
	title := c.GetString("title")
	tmpArticleTypeId := c.GetString("article_type")
	articleTypeId, err := strconv.ParseUint(tmpArticleTypeId, 10, 64)
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
	logs.Info(title, articleTypeId, content, author, fh.Size, fh.Filename, fh.Header, newName)
	c.TplName = "article.tpl"
	if title == "" || content == "" || author == "" {
		c.Data["message"] = "有字段为空"
		return
	}
	o := orm.NewOrmUsingDB("test")
	a := &models.Article{}
	a.Title = title
	a.Content = content
	// 查询频道对象，下拉框的处理
	at := &models.ArticleType{
		Id: articleTypeId,
	}
	err = o.Read(at)
	if err != nil {
		c.Data["message"] = "获取频道错误"
		return
	}
	a.ArticleType = at // 这里很重要
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
	c.Redirect("/article", 302)
	return
	c.Data["message"] = "写入文章成功，ID：" + strconv.FormatInt(id, 10)
}

// ArticleList 文章列表页
func (c *ArticleController) ArticleList() {
	msg := c.GetString("msg")
	tmpPage := c.GetString("page", "1")
	tmpType := c.GetString("type") // 频道ID
	page, _ := strconv.Atoi(tmpPage)
	articleTypeId, _ := strconv.Atoi(tmpType)
	pageSize := 3
	c.TplName = "article_list.tpl"
	o := orm.NewOrmUsingDB("test")
	// 获取频道列表
	var articleTypes []models.ArticleType
	_, err := o.QueryTable(models.ArticleType{}).All(&articleTypes)
	// 查询总数
	var articles []models.Article
	qs := o.QueryTable(models.Article{}).RelatedSel("ArticleType")
	if articleTypeId > 0 {
		qs = qs.Filter("ArticleType__Id", articleTypeId)
	}
	count, err := qs.Count()
	if err != nil {
		logs.Info(err)
		return
	}
	// 页数
	tmpPageCount := math.Ceil(float64(count) / float64(pageSize))
	pageCount := int(tmpPageCount)
	// 上下页处理
	if page > pageCount {
		page = pageCount
	}
	nextPage := page + 1
	if nextPage > pageCount {
		nextPage = pageCount
	}
	// 当前页数据
	start := (page - 1) * pageSize
	_, err = qs.Limit(pageSize, start).OrderBy("-id").All(&articles) // 关联表
	if err != nil {
		logs.Info(err)
		return
	}
	//logs.Info(articleTypesCount, count, num, pageCount)
	//logs.Info(articles)
	c.Data["articleTypes"] = articleTypes
	c.Data["type"] = articleTypeId
	c.Data["articles"] = articles
	c.Data["message"] = msg
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["page"] = page
	c.Data["nextPage"] = nextPage
}

// ArticleInfo 文章详情页
func (c *ArticleController) ArticleInfo() {
	c.Data["headTitle"] = "编辑文章页"
	c.Layout = "layout.tpl"
	c.TplName = "article_info.tpl"
	paramId := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		c.Redirect("/article_list", 302)
		return
	}
	// ID非法
	if id < 1 {
		c.Redirect("/article_list", 302)
		return
	}
	o := orm.NewOrmUsingDB("test")
	a := &models.Article{
		Id: id,
	}
	err = o.Read(a)
	if err == orm.ErrNoRows {
		c.Redirect("/article_list", 302)
		return
	} else if err != nil {
		c.Redirect("/article_list", 302)
		return
	}
	logs.Info(a)
	a.Pv += 1
	o.Update(a)
	// 多对多插入读者
	// 获取多对多操作对象
	m2m := o.QueryM2M(a, "Users")
	// 获取插入对象
	uid, _ := c.GetSession("beego_uid").(int)
	u := &models.User{
		Id: int64(uid),
	}
	o.Read(u)
	// 多对多插入
	_, err = m2m.Add(u)
	if err != nil {
		logs.Info(err)
		return
	}
	c.Data["article"] = a
}

// ArticleUpdate 文章内容编辑
func (c *ArticleController) ArticleUpdate() {
	tmpId := c.GetString("id")
	id, _ := strconv.ParseUint(tmpId, 10, 64)
	title := c.GetString("title")
	content := c.GetString("content")
	author := c.GetString("author")
	logs.Info(id, title, content, author)
	o := orm.NewOrmUsingDB("test")
	a := &models.Article{
		Id: id,
	}
	a.Title = title
	a.Content = content
	a.Author = author
	a.UpdateTime = time.Now()
	num, err := o.Update(a, "Title", "Content", "Author", "UpdateTime")
	logs.Info(num, err)
	//c.Data["message"] = "更新成功"
	//err = o.Read(a)
	//c.Data["article"] = a
	c.Redirect("/article_list", 302)
	return
}
func (c *ArticleController) ArticleDelete() {
	paramId := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseUint(paramId, 10, 64)
	o := orm.NewOrmUsingDB("test")
	a := &models.Article{
		Id: id,
	}
	err := o.Read(a)
	if err != nil {
		c.Redirect("/article_list?当前ID不存在", 302)
		return
	}
	num, err := o.Delete(a)
	logs.Info(num)
	if err != nil {
		c.Redirect("/article_list?msg=删除失败", 302)
		return
	}
	c.Redirect("/article_list?msg=删除成功", 302)
	return
}
