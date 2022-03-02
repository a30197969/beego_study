package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// 表的设计
type User struct {
	Id       int64     `orm:"auto"`
	Name     string    `orm:"size(50)"`
	Password string    `orm:"size(32)"`
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
func SelectUserList() []*User {
	o := orm.NewOrmUsingDB("test")
	qs := o.QueryTable(&User{})
	qs.Filter("name", "123")
	var users []*User
	num, err := qs.All(users)
	println(num, err)
	return users
}

// 根据主键ID查询
func SelectUser(id int64) (u User, err error) {
	o := orm.NewOrmUsingDB("test")
	u = User{}
	u.Id = id
	err = o.Read(&u)
	if err == orm.ErrNoRows {
		return u, nil
	} else if err == orm.ErrMissPK {
		return u, nil
	} else {
		return u, err
	}
	return u, nil
}
