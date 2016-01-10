package admin

import (
	"goblog/models"

	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Theme  string
	Option map[string]string
}

func (this *BaseController) Prepare() {
	// 检查登陆
	this.SetSession("username", "jiayx")
	this.SetSession("uid", 1)

	username := this.GetSession("username")
	if username == nil {
		this.Redirect("/admin/login", 302)
	}
	// 获取header中的分类
	categories, err := new(models.Category).All()
	if err != nil {
		this.Abort("500")
	}
	this.Data["Categories"] = categories

	// 配置项
	var option models.Option
	this.Option, _ = option.All()

	/*if _, ok := this.Option["theme"]; !ok {
		this.Option["theme"] = "default" // 博客前端默认主题
	}*/
	this.Theme = "admin"
}

func (this *BaseController) Display(path string) {
	this.Layout = this.Theme + "/layout.tpl"
	this.TplNames = this.Theme + "/" + path
}

func (this *BaseController) ShowMsg(msg, redirect string) {
	this.Data["msg"] = msg
	this.Data["redirect"] = redirect
	this.TplNames = this.Theme + "/show_msg.tpl"
	this.Render()
	this.StopRun()
}

func (this *BaseController) ArticleSaveRedirectUrl(id interface{}) string {
	url := "/admin/article/write"
	if id.(int64) > 0 {
		url = fmt.Sprintf("/admin/article/edit/%d", id.(int64))
	}
	return url
}

func (this *BaseController) SaySaveRedirectUrl(id interface{}) string {
	url := "/admin/say/write"
	if id.(int64) > 0 {
		url = fmt.Sprintf("/admin/say/edit/%d", id.(int64))
	}
	return url
}

func (this *BaseController) CategorySaveRedirectUrl(id interface{}) string {
	url := "/admin/category/write"
	if id.(int64) > 0 {
		url = fmt.Sprintf("/admin/category/edit/%d", id.(int64))
	}
	return url
}
