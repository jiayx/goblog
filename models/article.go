package models

import (
	// "fmt"
	"time"

	"goblog/helpers"

	"github.com/astaxie/beego/orm"
)

func AddArticle(title, content string, categories []string) (int64, error) {
	o := orm.NewOrm()
	uid := "1"
	userId, err := helpers.Str2Int(uid)

	article := &Article{
		Uid:        userId,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	var id int64
	id, err = o.Insert(article)

	// 写入文章分类关系表
	cateLen := len(categories)
	mapping := make([]CategoryArticleMap, cateLen)
	if cateLen > 0 {
		var cam CategoryArticleMap
		for i, c := range categories {
			cam.Aid = id
			cam.ShortName = c
			mapping[i] = cam
		}
	}

	_, err = o.InsertMulti(100, mapping)
	return id, err
}

// 1分类 2 排序方式
func GetArticles(args ...string) ([]*Article, error) {

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

	return articles, err
}

func GetArticleById(id string) (Article, error) {
	idInt, err := helpers.Str2Int(id)
	var article Article
	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("id", idInt)
	err = qs.One(&article)
	return article, err
}

func SetCategoryArticleMaping(maping []CategoryArticleMap) (int64, error) {
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, maping)
	return successNums, err
}
