package models

import (
	"goblog/helpers"

	// "fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id         int64
	Pid        int64 `orm:"default(0)"`
	Name       string
	ShortName  string
	Describe   string
	Sort       int64 `orm:"default(0)"`
	CreateTime time.Time
	UpdateTime time.Time
	Articles   []*Article `orm:"rel(m2m)"`
}

type CategoryRelation struct {
	*Category
	SubCategories []*CategoryRelation
}

// 获取分类名称
func (m *Category) GetCategoryNameByShortName(shortName string) (string, error) {
	var category Category
	o := orm.NewOrm()
	qs := o.QueryTable("category").Filter("ShortName", shortName)
	err := qs.One(&category, "name")
	return category.Name, err
}

func (m *Category) Insert() (int64, error) {

	id, err := orm.NewOrm().Insert(m)
	return id, err
}

func (m *Category) Update(fields ...string) error {
	o := orm.NewOrm()
	err := o.Read(&Category{Id: m.Id})
	if err == nil {
		_, err = o.Update(m, fields...)
	}
	return err
}

func (m *Category) Delete() error {
	_, err := orm.NewOrm().Delete(m)
	return err
}

// 获取全部分类 - 详情
func (m *Category) AllDetail() ([]*CategoryRelation, error) {
	categories := make([]*Category, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Category").OrderBy("Sort").All(&categories)
	for _, cate := range categories {
		o.LoadRelated(cate, "Articles")
	}
	_categories := BuildCategory(categories, 0)
	// fmt.Println(_categories)
	return _categories, err
}

// 获取全部分类 - 仅获取分类
func (m *Category) All() ([]*CategoryRelation, error) {
	categories := make([]*Category, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Category").OrderBy("Sort").All(&categories)
	_categories := BuildCategory(categories, 0)
	// fmt.Println(_categories)
	return _categories, err
}

func (m *Category) One(id string) (Category, error) {
	idInt := helpers.Str2Int(id)
	var category Category
	o := orm.NewOrm()
	qs := o.QueryTable("say").Filter("id", idInt)
	err := qs.One(&category)
	return category, err
}

func BuildCategory(categories []*Category, pid int64) []*CategoryRelation {
	_category := make([]*CategoryRelation, 0)
	for _, cate := range categories {
		if cate.Pid == pid {
			cateRel := new(CategoryRelation)
			// fmt.Println(cate)
			cateRel.Category = cate
			cateRel.SubCategories = BuildCategory(categories, cate.Id)
			_category = append(_category, cateRel)
		}
	}
	return _category
}
