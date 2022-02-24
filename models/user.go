package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int64     `orm:"auto"`
	Name     string    `orm:"size(50)"`
	Password string    `orm:"size(20)"`
	RegTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	// 将你定义的 Model 进行注册
	orm.RegisterModel(new(User))
	//orm.NewOrmUsingDB("test")
}

// InsertUser 注册用户
func InsertUser(name string, password string) (id int64, err error) {
	u := &User{
		Name:     name,
		Password: password,
	}
	o := orm.NewOrmUsingDB("test")
	id, err = o.Insert(u)
	return
}
