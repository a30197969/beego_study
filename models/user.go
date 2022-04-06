package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// User 表的设计
type User struct {
	Id       int64      `orm:"auto"`
	Name     string     `orm:"size(50);unique"`
	Password string     `orm:"size(32)"`
	RegTime  time.Time  `orm:"auto_now_add;type(datetime)"`
	Articles []*Article `orm:"reverse(many)"`
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

// SelectUserById 根据ID，查询一个注册用户
func SelectUserById(id int64) (u *User, err error) {
	o := orm.NewOrmUsingDB("test")
	u = &User{
		Id: id,
	}
	err = o.Read(u)
	if err != nil {
		logs.Warning(err)
	}
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err == orm.ErrMissPK {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return u, nil
}

// SelectUserList 查询所有的注册用户
func SelectUserList() ([]*User, error) {
	o := orm.NewOrmUsingDB("test")
	var users []*User
	num, err := o.QueryTable("user").Filter("name__icontains", "pengliang").All(&users) // 一定要用指针
	logs.Info(num, err)
	return users, err
}
