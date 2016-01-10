package models

import (
	"goblog/helpers"

	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int64
	Title      string
	Content    string `orm:"size(5000)"`
	Views      int64  `orm:"index"`
	Status     int8   `orm:"default(1)"` // 0删除 1正常 2草稿
	CreateTime time.Time
	UpdateTime time.Time
	Categories []*Category `orm:"reverse(many)"`
	User       *User       `orm:"rel(fk)"`
	Comments   []*Comment  `orm:"reverse(many)"`
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Article) Insert() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	m.Id = id
	if len(m.Categories) > 0 {
		o.QueryM2M(m, "Categories").Add(m.Categories)
	}
	return id, err
}

func (m *Article) Update(fields ...string) error {
	o := orm.NewOrm()
	err := o.Read(&Article{Id: m.Id})
	if err == nil {
		_, err = o.Update(m, fields...)
		o.QueryM2M(m, "Categories").Clear()
		if len(m.Categories) > 0 {
			o.QueryM2M(m, "Categories").Add(m.Categories)
		}
	}
	return err
}

func (m *Article) Delete() error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.QueryM2M(m, "Categories").Clear()
	if err == nil {
		_, err = o.Delete(m)
	}
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

// 1分类 2 排序方式
func (m *Article) All(args ...string) ([]*Article, error) {

	cate := ""
	orderBy := "-CreateTime"
	if len(args) == 2 {
		cate = args[0]
		orderBy = args[1]
	} else if len(args) == 1 {
		cate = args[0]
	}

	articles := make([]*Article, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("article")
	if cate != "" {
		qs = qs.Filter("category", cate)
	}

	qs = qs.OrderBy(orderBy)
	_, err := qs.All(&articles)

	for _, a := range articles {
		o.LoadRelated(a, "Categories")

	}
	return articles, err
}

func (m *Article) One(id string, isFront bool) (Article, error) {
	idInt := helpers.Str2Int(id)
	var article Article
	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("id", idInt)
	err := qs.One(&article)
	o.LoadRelated(&article, "Categories")
	fmt.Println(article)
	if isFront {
		article.ViewAdd()
	}
	return article, err
}

func (m *Article) ViewAdd(args ...int64) {
	var count int64 = 1
	if len(args) > 0 {
		count = args[0]
	}
	m.Views = m.Views + count
	o := orm.NewOrm()
	o.Update(m, "Views")
}
