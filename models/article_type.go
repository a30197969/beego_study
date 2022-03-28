package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type ArticleType struct {
	Id       uint64     `orm:"auto"`
	Name     string     `orm:"size(100)"`
	AddTime  time.Time  `orm:"auto_now_add;type(datetime)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	// 将你定义的 Model 进行注册
	orm.RegisterModel(new(ArticleType))
	//orm.NewOrmUsingDB("test")
}
