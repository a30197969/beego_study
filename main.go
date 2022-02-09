package main

import (
	_ "beego_study/routers" // 只是调用init方法
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/forum?charset=utf8mb4&parseTime=true&loc=Local")
	orm.RegisterDataBase("test", "mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
}

func main() {
	// 静态文件处理
	beego.SetStaticPath("/js", "static/js")
	beego.Run()
}
