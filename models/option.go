package models

import (
	// "fmt"
	// "time"

	"github.com/astaxie/beego/orm"
)

// 系统配置表
type Option struct {
	Name  string `orm:"pk"`
	Value string
}

func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Option) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(m)
	return id, err
}

func (m *Option) Update(fields ...string) error {
	o := orm.NewOrm()
	err := o.Read(&Option{Name: m.Name})
	if err == nil {
		_, err = o.Update(m, fields...)
	}
	return err
}

func (m *Option) Delete() error {

	_, err := orm.NewOrm().Delete(m)
	return err
}

// 1分类 2 排序方式
func (m *Option) All() (map[string]string, error) {

	options := make([]*Option, 0)
	_, err := orm.NewOrm().QueryTable("option").All(&options)

	optionMap := make(map[string]string)
	optionMap["sayAllowComment"] = "1"     // 0不可评论 1可评论 2需审核
	optionMap["articleAllowComment"] = "1" // 0不可评论 1可评论 2需审核
	optionMap["commentsPageSize"] = "10"   // 每页显示的评论个数
	optionMap["indexPageSize"] = "10"      // 首页展示的文章数
	optionMap["commentSortType"] = "DESC"  //评论排序
	optionMap["theme"] = "default"         //评论排序
	return optionMap, err
}

func (m *Option) One(name string) (string, error) {

	var option Option
	o := orm.NewOrm()
	qs := o.QueryTable("option").Filter("name", name)
	err := qs.One(&option)
	return option.Name, err
}
