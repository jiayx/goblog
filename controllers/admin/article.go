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

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["MdStyles"] = "admin/md_css.tpl"
	this.LayoutSections["MdScripts"] = "admin/md_js.tpl"

	this.Display("article_add.tpl")
}

// 保存
func (this *ArticleController) Save() {

	var (
		err     error
		article models.Article
		// categoryArticles models.CategoryArticles
	)

	id, _ := this.GetInt64("id", 0)
	uid := this.GetSession("uid")
	fmt.Println(article)

	var user models.User
	user.Id = int64(uid.(int))
	article.User = &user
	article.Title = strings.TrimSpace(this.GetString("title", "未命名"))
	article.Content = strings.TrimSpace(this.GetString("content"))
	article.Views = 0
	article.Status, err = this.GetInt8("status", 1)
	article.UpdateTime = time.Now()
	fmt.Println(article.UpdateTime)
	categories := this.GetStrings("categories[]")

	if article.Content == "" {
		this.ShowMsg("内容不能为空", this.ArticleSaveRedirectUrl(id))
	}

	cateLen := len(categories)
	mapping := make([]*models.Category, 0)
	if cateLen > 0 {
		for _, c := range categories {
			var ca models.Category
			ca.Id = helpers.Str2Int(c)
			mapping = append(mapping, &ca)
		}
	}
	article.Categories = mapping

	if id > 0 {
		article.Id = id
		err = article.Update("Title", "Content", "UpdateTime")

	} else {
		article.CreateTime = time.Now()
		id, err = article.Insert()
	}
	fmt.Println(err, "sdsggrerrr")
	if err != nil {
		this.ShowMsg(err.Error(), this.ArticleSaveRedirectUrl(id))
	}

	this.Redirect("/admin/manage/post", 302)
}

// 列表
func (this *ArticleController) List() {
	var article models.Article
	list, err := article.All(map[string]string{})
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
	article, err := article.One(id, false)
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/post")
	}
	// fmt.Println(article.Categories)
	this.Data["Title"] = "编辑文章"
	this.Data["Article"] = article
	this.Data["Cates"] = article.Categories

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["MdStyles"] = "admin/md_css.tpl"
	this.LayoutSections["MdScripts"] = "admin/md_js.tpl"

	this.Display("article_edit.tpl")
}

// 删除
func (this *ArticleController) Delete() {

	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("404")
	}
	idInt := helpers.Str2Int(id)
	article := models.Article{Id: idInt}
	err := article.Delete()
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/post")
	}
	this.Redirect("/admin/manage/post", 302)
}
