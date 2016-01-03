package admin

import (
	"goblog/helpers"
	"goblog/models"

	"fmt"
	"strings"
	"time"
)

type ArticleController struct {
	BaseController
}

// 新增
func (this *ArticleController) Write() {
	this.Data["Title"] = "写文章"
	this.Display("article_add.tpl")
}

// 保存
func (this *ArticleController) Save() {

	var (
		err              error
		article          models.Article
		categoryArticles models.CategoryArticles
	)

	id, _ := this.GetInt64("id", 0)
	uid := this.GetSession("uid")
	article.Uid = int64(uid.(int))
	article.Title = strings.TrimSpace(this.GetString("title", "未命名"))
	article.Content = strings.TrimSpace(this.GetString("content"))
	article.Views = 0
	article.Status, err = this.GetInt8("status", 1)
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()

	categories := this.GetStrings("categories[]")

	if article.Content == "" {
		// 显示错误 、 返回json
		this.ShowMsg("内容不能为空", "/admin/article/write")
	}

	cateLen := len(categories)
	mapping := make([]models.CategoryArticles, 0)

	if id > 0 {
		article.Id = id
		_, err = article.Update()

	} else {
		id, err = article.Insert()
	}

	if err != nil {
		this.ShowMsg("出错了请重试", "/admin/article/write")
	}

	if cateLen > 0 {
		var ca models.CategoryArticles
		for _, c := range categories {
			ca.Aid = id
			ca.Cid = helpers.Str2Int(c)
			mapping = append(mapping, ca)
		}
	}

	categoryArticles.UpdateMapping(id, &mapping)
	fmt.Println(mapping)
	if err != nil {
		this.ShowMsg("出错了请重试", "/admin/article/write")
	}

	this.Redirect("/admin/manage/post", 302)
}

// 列表
func (this *ArticleController) List() {
	var article models.Article
	list, err := article.GetArticles()
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/post")
	}
	this.Data["Title"] = "管理文章"
	this.Data["List"] = list
	this.Display("article_list.tpl")
}

// 编辑
func (this *ArticleController) Edit() {

	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("404")
	}
	var article models.Article
	post, err := article.GetArticleById(id)
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/post")
	}
	this.Data["Title"] = "编辑文章"
	this.Data["Article"] = post
	this.Display("article_edit.tpl")
}
