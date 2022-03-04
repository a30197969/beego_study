package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Article struct {
	Id          uint64 `orm:"auto"`
	Title       string `orm:"size(100)"`
	Content     string `orm:"type(text);null"`
	Pv          uint
	QiniuBucket string    `orm:"size(50)"`
	QiniuKey    string    `orm:"size(255)"`
	Author      string    `orm:"size(50)"`
	AddTime     time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime);index"`
}

func init() {
	// 将你定义的 Model 进行注册
	orm.RegisterModel(new(Article))

	//orm.NewOrmUsingDB("test")
}
