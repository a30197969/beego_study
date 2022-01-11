package controllers

import (
	"beego_study/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/pkg/errors"
)

type SubjectController struct {
	beego.Controller
}

func (c *SubjectController) Get() {
	var subject models.Subject
	// 匿名函数，校验相关的代码模块化处理，代码的可读性、可维护性高
	err := func() error {
		id, err := c.GetInt("id")
		//beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()
	if err != nil {
		c.Ctx.WriteString("wrong params")
		return
	}
	var option map[string]string
	if err = json.Unmarshal([]byte(subject.Option), &option); err != nil {
		c.Ctx.WriteString("wrong params, json decode")
		return
	}
	c.Data["Id"] = subject.Id
	c.Data["Option"] = option
	c.Data["Img"] = "https://image.xcar.com.cn/attachments/a/day_220110/2022011014_28e7bf645b8b7e3dc4d8mJw1F1z2i2Nd.jpg?imageView2/1/w/362/h/240"
	c.TplName = "subject.tpl"
}
func (c *SubjectController) Post() {
	var subject models.Subject
	// 匿名函数，校验相关的代码模块化处理，代码的可读性、可维护性高
	err := func() error {
		id, err := c.GetInt("id")
		//beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()
	if err != nil {
		c.Ctx.WriteString("wrong params")
	}
	answer := c.GetString("key")
	right := models.Answer(subject.Id, answer)
	c.Data["Right"] = right
	c.Data["Next"] = subject.Id + 1
	c.Data["Id"] = subject.Id
	c.TplName = "subject.tpl"
}
