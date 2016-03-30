package controllers

import (
	"goblog/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Theme  string
	Option map[string]string
}

func (this *BaseController) Prepare() {
	// 获取header中的分类
	categories, err := new(models.Category).All()
	if err != nil {
		this.Abort("500")
	}
	this.Data["Categories"] = categories

	// 配置项
	var option models.Option
	this.Option, _ = option.All()

	if _, ok := this.Option["theme"]; !ok {
		this.Option["theme"] = "default" // 博客前端默认主题
		this.Theme = "admin"
	} else {
		this.Theme = this.Option["theme"]
	}

}

func (this *BaseController) Display(path string) {
	// this.Layout = this.Theme + "/layout.tpl"
	this.TplName = this.Theme + "/" + path
}

func (this *BaseController) ShowMsg(msg, redirect string) {
	this.Data["msg"] = msg
	this.Data["redirect"] = redirect
	this.TplName = this.Theme + "/show_msg.tpl"
	this.Render()
	this.StopRun()
}
