package admin

import (
	"goblog/models"
)

type ArticleController struct {
	BaseController
}

/*func (this *ArticleController) Get() {
	id := this.Ctx.Input.Param(":id")
	fmt.Println(id)
	article, err := models.GetArticleById(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["Title"] = article.Title
	this.Data["Article"] = article
	this.TplNames = "viewArticle.tpl"
}

func (this *ArticleController) Post() {
	title := this.GetString("title")
	content := this.GetString("content")
	if title == "" {
		// 显示错误 、 返回json
		this.Abort("401")
	}
	models.AddArticle(title, content, "php")
	this.Redirect("/", 302)
}*/

// 新增
func (this *ArticleController) Write() {
	this.Data["Title"] = "写文章"
	this.Display("article_add.tpl")
}

// 保存
func (this *ArticleController) Save() {
	title := this.GetString("title", "未命名")
	content := this.GetString("content")
	categories := this.GetStrings("categories[]")

	if content == "" {
		// 显示错误 、 返回json
		this.ShowMsg("内容不能为空", "/admin/article/write")
	}
	_, err := models.AddArticle(title, content, categories)
	if err != nil {
		this.ShowMsg("出错了请重试", "/admin/article/write")
	}
	this.Redirect("/admin/manage/post", 302)
}

// 列表
func (this *ArticleController) List() {

	list, err := models.GetArticles()
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
	article, err := models.GetArticleById(id)
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/post")
	}
	this.Data["Title"] = "编辑文章"
	this.Data["Article"] = article
	this.Display("article_edit.tpl")
}
