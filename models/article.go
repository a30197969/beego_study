package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Article struct {
	Id          uint64       `orm:"auto"`
	Title       string       `orm:"size(100)"`
	Content     string       `orm:"type(text);null"`
	ArticleType *ArticleType `orm:"rel(fk)"` // 设置外键
	Pv          uint
	Users       []*User   `orm:"rel(m2m)"` // 多对多关系
	QiniuBucket string    `orm:"size(100)"`
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
