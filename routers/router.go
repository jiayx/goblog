package routers

import (
	"goblog/controllers"
	"goblog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	// 前端留言板
	beego.Router("/", &controllers.IndexController{})

	// 后台管理
	beego.Router("/admin", &admin.IndexController{})
	beego.Router("/admin/login", &admin.MemberController{})
	beego.Router("/admin/logout", &admin.MemberController{}, "get:Logout")

	// 文章
	beego.Router("/admin/article/write", &admin.ArticleController{}, "get:Write")
	beego.Router("/admin/article/edit/:id:int", &admin.ArticleController{}, "get:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete/:id:int", &admin.ArticleController{}, "get:Delete")
	beego.Router("/admin/manage/post", &admin.ArticleController{}, "get:List")

	// 说说
	beego.Router("/admin/say/write", &admin.SayController{}, "get:Write")
	beego.Router("/admin/say/edit/:id:int", &admin.SayController{}, "get:Edit")
	beego.Router("/admin/say/save", &admin.SayController{}, "post:Save")
	beego.Router("/admin/say/delete/:id:int", &admin.SayController{}, "get:Delete")
	beego.Router("/admin/manage/say", &admin.SayController{}, "get:List")

	// 分类
	beego.Router("/admin/manage/category", &admin.CategoryController{}, "get:List")

}
