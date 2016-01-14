package models

import (
	"goblog/helpers"

	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Say struct {
	Id         int64
	Uid        int64
	Content    string `orm:"size(2000)"`
	Views      int64  `orm:"index"`
	Status     int8   `orm:"default(1)"` // 0删除 1正常 2草稿
	CreateTime time.Time
	UpdateTime time.Time
}

func (m *Say) Insert() (int64, error) {

	id, err := orm.NewOrm().Insert(m)
	return id, err
}

func (m *Say) Update(fields ...string) error {
	o := orm.NewOrm()
	err := o.Read(&Say{Id: m.Id})
	if err == nil {
		_, err = o.Update(m, fields...)
	}
	return err
}

func (m *Say) Delete() error {
	_, err := orm.NewOrm().Delete(m)
	return err
}

// 1分类 2 排序方式
func (m *Say) All(args ...string) ([]*Say, error) {

	var status int8 = 1
	orderBy := "-CreateTime"
	if len(args) == 2 {
		status = helpers.Str2Int8(args[0])
		orderBy = args[1]
	} else if len(args) == 1 {
		status = helpers.Str2Int8(args[0])
	}

	says := make([]*Say, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("say")

	qs = qs.Filter("Status", status)
	qs = qs.OrderBy(orderBy)
	_, err := qs.All(&says)

	return says, err
}

func (m *Say) One(id string, isFront bool) (Say, error) {
	idInt := helpers.Str2Int(id)
	var say Say
	o := orm.NewOrm()
	qs := o.QueryTable("say").Filter("id", idInt)
	err := qs.One(&say)
	fmt.Println(say)
	if isFront {
		say.ViewAdd()
	}
	return say, err
}

func (m *Say) ViewAdd(args ...int64) {
	var count int64 = 1
	if len(args) > 0 {
		count = args[0]
	}
	m.Views = m.Views + count
	o := orm.NewOrm()
	o.Update(m, "Views")
}

// 总数
func (m *Say) Count() int64 {
	count, _ := orm.NewOrm().QueryTable("say").Count()
	return count
}
