package admin

import (
	"goblog/helpers"
	"goblog/models"

	"fmt"
	"strings"
	"time"
)

type CategoryController struct {
	BaseController
}

// 新增
func (this *CategoryController) Write() {
	this.Data["Title"] = "写说说"
	this.Display("category_add.tpl")
}

// 保存
func (this *CategoryController) Save() {

	var (
		err      error
		category models.Category
	)

	id, _ := this.GetInt64("id", 0)
	category.Pid, _ = this.GetInt64("id", 0)
	category.Name = strings.TrimSpace(this.GetString("name"))
	category.ShortName = strings.TrimSpace(this.GetString("short_name"))
	category.Describe = strings.TrimSpace(this.GetString("describe"))
	category.Sort, _ = this.GetInt64("id", 0)
	category.UpdateTime = time.Now()

	if category.Name == "" {
		this.ShowMsg("名称不能为空", this.CategorySaveRedirectUrl(id))
	}

	if id > 0 {
		category.Id = id
		err = category.Update("Content", "UpdateTime")

	} else {
		category.CreateTime = time.Now()
		id, err = category.Insert()
	}
	fmt.Println(err, "sdsggrerrr")
	if err != nil {
		this.ShowMsg(err.Error(), this.CategorySaveRedirectUrl(id))
	}

	this.Redirect("/admin/manage/category", 302)
}

// 列表
func (this *CategoryController) List() {
	var category models.Category
	list, err := category.All()
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/category")
	}
	this.Data["Title"] = "分类管理"
	this.Data["List"] = list
	this.Display("category_list.tpl")
}

// 编辑
func (this *CategoryController) Edit() {

	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("404")
	}
	var category models.Category
	category, err := category.One(id)
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/category")
	}
	// fmt.Println(category.Categories)
	this.Data["Title"] = "编辑说说"
	this.Data["Category"] = category
	this.Display("category_edit.tpl")
}

// 删除
func (this *CategoryController) Delete() {

	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("404")
	}
	idInt := helpers.Str2Int(id)
	category := models.Category{Id: idInt}
	err := category.Delete()
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/category")
	}
	this.Redirect("/admin/manage/category", 302)
}
