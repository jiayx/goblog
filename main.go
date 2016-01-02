package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goblog/models"
	_ "goblog/routers"
)

func init() {
	// 注册数据库
	models.Init()
}

func main() {
	// beego.EnableAdmin = true
	// 开启 ORM 调试模式
	// orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.AddFuncMap("InArray", InArray)
	beego.Run()
}

func InArray(val string, array []*models.Category) bool {
	for _, value := range array {
		if value.ShortName == val {
			return true
		}
	}
	return false
}
