package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id         int64
	Uid        int64
	Title      string
	Content    string `orm:"size(5000)"`
	Views      int64  `orm:"index"`
	Status     int8   `oem:"default(1)"` // 0删除 1正常 2草稿
	CreateTime time.Time
	UpdateTime time.Time
}

type Category struct {
	Id         int64
	Pid        int64 `orm:"default(0)"`
	Name       string
	ShortName  string
	Describe   string
	Sort       int64 `orm:"default(0)"`
	CreateTime time.Time
	UpdateTime time.Time
}

type User struct {
	Id         int64
	Username   string
	Password   string `orm:"default(32)"`
	CreateTime time.Time
	UpdateTime time.Time
	IsDeleted  int `orm:"default(0)"`
}

// 分类与文章对应关系
type CategoryArticleMap struct {
	Id        int64
	ShortName string `orm:"index"`
	Aid       int64  `orm:"index"`
}

func Init() {
	orm.RegisterModel(new(Article), new(Category), new(User), new(CategoryArticleMap))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/goblog?charset=utf8", 30)

	beego.AddFuncMap("InCategoryArray", InCategoryArray)

}

func CreateUser(username, password string) (int64, error) {
	o := orm.NewOrm()

	User := &User{
		Username:   username,
		Password:   password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	id, err := o.Insert(User)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindUser(i interface{}) (User, error) {

	var user User
	o := orm.NewOrm()
	qs := o.QueryTable("user")

	switch i.(type) {
	case string:
		qs = qs.Filter("username", i)
	case int64:
		qs = qs.Filter("id", i)
	}
	err := qs.One(&user)
	return user, err
}

func InCategoryArray(val string, array []*Category) bool {
	for _, value := range array {
		if value.ShortName == val {
			return true
		}
	}
	return false
}
