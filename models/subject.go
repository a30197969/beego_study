package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

type Subject struct {
	Id        int    `orm:"auto"`
	Option    string `orm:"size(255)"`
	AnswerKey string `orm:"column(answer_key)"`
	Status    int8
	Img       string
	Dateline  time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Subject))
}
func GetSubject(id int) (s Subject, err error) {
	o := orm.NewOrmUsingDB("test")
	s = Subject{Id: id}
	err = o.Read(&s)
	return
}
func Answer(sid int, answerkey string) bool {
	subject, err := GetSubject(sid)
	if err != nil {
		return false
	}
	return strings.Compare(strings.ToUpper(answerkey), subject.AnswerKey) == 0
}
