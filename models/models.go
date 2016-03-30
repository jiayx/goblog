package models

import (
	// "fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int64
	Username   string
	Password   string `orm:"size(32)"`
	CreateTime time.Time
	UpdateTime time.Time
	IsDeleted  int        `orm:"default(0)"`
	Articles   []*Article `orm:"reverse(many)"`
}

type Comment struct {
	Id         int64
	ReplyId    int64     // 回复的留言id 为0是回复文章
	Name       string    `orm:"null"`
	Ip         string    `orm:"null"`
	CreateTime time.Time `orm:"null"`
	Say        *Say      `orm:"rel(fk);default(0)"`
	Article    *Article  `orm:"rel(fk);default(0)"`
	IsDeleted  int       `orm:"default(0)"`
}

func init() {
	// 初始化数据库
	orm.RegisterModel(new(Article), new(Category), new(User), new(Say), new(Option), new(Comment))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/goblog?charset=utf8&loc=Asia%2FShanghai", 30)

	// 注册模板函数
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

func InCategoryArray(val int64, array []*Category) bool {
	// fmt.Println(val, array)
	for _, value := range array {
		if value.Id == val {
			return true
		}
	}
	return false
}
