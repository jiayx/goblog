package admin

import (
	"goblog/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	article := new(models.Article)
	latestArticles, err := article.All(map[string]string{"pageSize": "5"})
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/post")
	}
	this.Data["ArticleCount"] = article.Count()
	this.Data["SayCount"] = new(models.Say).Count()
	this.Data["CategoryCount"] = new(models.Category).Count()
	this.Data["LatestArticles"] = latestArticles
	this.Display("index.tpl")
}
