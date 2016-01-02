package admin

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type BaseController struct {
	beego.Controller
	Theme string
}

func (this *BaseController) Prepare() {
	// 检查登陆
	/*username := this.GetSession("username")
	if username == nil {
		this.Redirect("/admin/login", 302)
	}*/
	// 获取header中的分类
	categories, err := models.GetCategories()
	if err != nil {
		this.Abort("500")
	}
	this.Data["Categories"] = categories

	//
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
