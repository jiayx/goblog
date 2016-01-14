package models

import (
	"goblog/helpers"

	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const PAGE_SIZE = 20

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
func (m *Article) All(queries map[string]string) ([]*Article, error) {

	articles := make([]*Article, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("article")

	// 类型筛选
	if category, ok := queries["category"]; ok {
		qs = qs.Filter("category", category)
	}
	// 排序方式
	if orderBy, ok := queries["orderBy"]; ok {
		qs = qs.OrderBy(orderBy)
	} else {
		qs = qs.OrderBy("-CreateTime")
	}
	// 分页
	if pageIndex, ok := queries["pageIndex"]; ok {
		pageIndexInt := helpers.Str2Int(pageIndex)
		if pageIndexInt < 1 {
			pageIndexInt = 1
		}
		qs = qs.Offset((pageIndexInt - 1) * PAGE_SIZE)
	} else {
		qs = qs.OrderBy("-CreateTime")
	}
	if pageSize, ok := queries["pageSize"]; ok {
		pageSizeInt := helpers.Str2Int(pageSize)
		qs = qs.Limit(pageSizeInt)
	} else {
		qs = qs.Limit(PAGE_SIZE)
	}

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

// 总数
func (m *Article) Count() int64 {
	count, _ := orm.NewOrm().QueryTable("article").Count()
	return count
}
