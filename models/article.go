package models

import (
	"goblog/helpers"

	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
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

// 分类与文章对应关系
type CategoryArticles struct {
	Id  int64
	Cid int64
	Aid int64
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Article) Insert() (int64, error) {

	id, err := orm.NewOrm().Insert(m)
	return id, err
}

func (m *Article) Update() (int64, error) {

	num, err := orm.NewOrm().Update(m)
	return num, err
}

func (m *CategoryArticles) UpdateMapping(id int64, mapping *[]CategoryArticles) {
	fmt.Println(id, mapping)
	o := orm.NewOrm()
	o.QueryTable(&CategoryArticles{}).Filter("Aid", id).Delete()
	o.InsertMulti(100, mapping)
}

// 1分类 2 排序方式
func (m *Article) GetArticles(args ...string) ([]*Article, error) {

	cate := ""
	orderBy := "-UpdateTime"
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

	categories := make([]Category, 0)
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("ca.id",
		"c.name", "c.ShortName").
		From("CategoryArticles as ca").
		InnerJoin("Category as c").On("ca.cid = c.id").
		Where("ca.aid = 1").
		Limit(10).Offset(0)
	sql := qb.String()
	/*o.QueryTable(&Category{}).Filter("Aid", id).All(Categories)*/
	o.Raw(sql, 20).QueryRows(&categories)

	return articles, err
}

func (m *Article) GetArticleById(id string) (Article, error) {
	idInt := helpers.Str2Int(id)
	var article Article
	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("id", idInt)
	err := qs.One(&article)
	return article, err
}

/*func SetCategoryArticleMaping(maping []CategoryArticles) (int64, error) {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, maping)
	return successNums, err
}*/
