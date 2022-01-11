package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type Thread struct {
	Threadid     int64  `orm:"pk;auto"`
	Title        string `orm:"size(250);column(title)"` // 设置varchar长度为250 列名为title
	Firstpostid  int
	Lastpost     int
	Forumid      int
	Pollid       int
	Open         int
	Replycount   int
	Hiddencount  int
	Postusername string `orm:"size(128)"`
	Postuserid   int
	Lastposter   int
	Dateline     string `orm:"size(128)"`
	Views        int
	Iconid       int
	Notes        string `orm:"size(128)"`
	Visible      int
	Sticky       int
	Globalsticky int
	Votenum      int
	Votetotal    int
	Attach       int
	Similar      string `orm:"size(128)"`
	Titlecolor   string `orm:"size(128)"`
	Titlebold    int
	Execposts    int
	Excerption   int
	FnPass       int
	Attachmentid int
	IsLock       int
	IsReview     int
	Flash        string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Thread))
}

// AddThread insert a new Thread into database and returns
// last inserted Id on success.
func AddThread(m *Thread) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetThreadById retrieves Thread by Id. Returns error if
// Id doesn't exist
func GetThreadById(id int64) (v *Thread, err error) {
	o := orm.NewOrm()
	v = &Thread{Threadid: id}
	err = o.Raw("select * from thread where threadid=? limit 1", id).QueryRow(v)
	if err != nil {
		return nil, err
	}
	return v, nil

	/*if err = o.QueryTable(new(Thread)).Filter("threadid", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err*/
}

// GetAllThread retrieves all Thread matches certain condition. Returns empty list if
// no records exist
func GetAllThread(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Thread))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Thread
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateThread updates Thread by Id and returns error if
// the record to be updated doesn't exist
func UpdateThreadById(m *Thread) (err error) {
	o := orm.NewOrm()
	v := Thread{Threadid: m.Threadid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteThread deletes Thread by Id and returns error if
// the record to be deleted doesn't exist
func DeleteThread(id int64) (err error) {
	o := orm.NewOrm()
	v := Thread{Threadid: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Thread{Threadid: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
