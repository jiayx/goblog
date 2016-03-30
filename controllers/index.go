package controllers

import (
// "goblog/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *IndexController) Index() {

	this.Display("index.tpl")
}
