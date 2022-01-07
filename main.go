package main

import (
	_ "beego_study/routers" // 只是调用init方法
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/forum?charset=utf8mb4&parseTime=true&loc=Local")
	beego.Run()
}
