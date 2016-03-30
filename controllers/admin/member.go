package admin

import (
	"github.com/astaxie/beego"
)

type MemberController struct {
	beego.Controller
}

func (this *MemberController) Get() {
	this.TplName = "admin/login.tpl"
}

func (this *MemberController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "" || password == "" {
		this.Redirect("/", 302)
	}
	// isAutoLogin, _ := this.GetBool("auto-login")

	if username == "jiayx" {
		this.SetSession("username", username)
		this.SetSession("uid", 1)
	}
	this.Redirect("/admin", 302)
}

func (this *MemberController) Logout() {

	this.SetSession("username", nil)
	this.Redirect("/admin/login", 302)
}

func (this *MemberController) Add() {

}
