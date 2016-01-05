package admin

import (
	"goblog/helpers"
	"goblog/models"

	"fmt"
	"strings"
	"time"
)

type SayController struct {
	BaseController
}

// 新增
func (this *SayController) Write() {
	this.Data["Title"] = "写说说"
	this.Display("say_add.tpl")
}

// 保存
func (this *SayController) Save() {

	var (
		err error
		say models.Say
	)

	id, _ := this.GetInt64("id", 0)
	uid := this.GetSession("uid")
	say.Uid = int64(uid.(int))
	say.Content = strings.TrimSpace(this.GetString("content"))
	say.Views = 0
	say.Status, err = this.GetInt8("status", 1)
	say.UpdateTime = time.Now()
	fmt.Println(say.UpdateTime)

	if say.Content == "" {
		this.ShowMsg("内容不能为空", this.SaySaveRedirectUrl(id))
	}

	if id > 0 {
		say.Id = id
		err = say.Update("Content", "UpdateTime")

	} else {
		say.CreateTime = time.Now()
		id, err = say.Insert()
	}
	fmt.Println(err, "sdsggrerrr")
	if err != nil {
		this.ShowMsg(err.Error(), this.SaySaveRedirectUrl(id))
	}

	this.Redirect("/admin/manage/say", 302)
}

// 列表
func (this *SayController) List() {
	var say models.Say
	list, err := say.All()
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/say")
	}
	this.Data["Title"] = "管理文章"
	this.Data["List"] = list
	this.Display("say_list.tpl")
}

// 编辑
func (this *SayController) Edit() {

	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("404")
	}
	var say models.Say
	say, err := say.One(id, false)
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/say")
	}
	// fmt.Println(say.Categories)
	this.Data["Title"] = "编辑说说"
	this.Data["Say"] = say
	this.Display("say_edit.tpl")
}

// 删除
func (this *SayController) Delete() {

	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("404")
	}
	idInt := helpers.Str2Int(id)
	say := models.Say{Id: idInt}
	err := say.Delete()
	if err != nil {
		this.ShowMsg("出错了", "/admin/manage/say")
	}
	this.Redirect("/admin/manage/say", 302)
}
