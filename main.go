package main

import (
	"goblog/models"
	_ "goblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 数据库初始化
	models.Init()

}

func main() {
	// beego.EnableAdmin = true
	// 开启 ORM 调试模式
	// orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
