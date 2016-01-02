package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 获取全部分类
func GetCategories() ([]*Category, error) {
	categories := make([]*Category, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("Category").OrderBy("Sort").All(&categories)
	return categories, err
}

// 获取分类名称
func GetCategoryNameByShortName(shortName string) (string, error) {
	var category Category
	o := orm.NewOrm()
	qs := o.QueryTable("category").Filter("ShortName", shortName)
	err := qs.One(&category, "name")
	return category.Name, err
}

func AddCategory(pid int64, name, shortName, describe string) error {
	o := orm.NewOrm()

	Category := &Category{
		Pid:        pid,
		Name:       name,
		ShortName:  shortName,
		Describe:   describe,
		Sort:       0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	_, err := o.Insert(Category)
	if err != nil {
		return err
	}
	return nil
}
