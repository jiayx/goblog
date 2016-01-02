package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers/admin"
)

func init() {

	// 后台管理
	beego.Router("/admin", &admin.IndexController{})
	beego.Router("/admin/login", &admin.MemberController{})
	beego.Router("/admin/logout", &admin.MemberController{}, "get:Logout")

	beego.Router("/admin/article/write", &admin.ArticleController{}, "get:Write")
	beego.Router("/admin/article/edit/:id:int", &admin.ArticleController{}, "get:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/manage/post", &admin.ArticleController{}, "get:List")

}
