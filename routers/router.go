package routers

import (
	"beego_study/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 路由配置 https://www.jianshu.com/p/417fcc5841f8
	//beego.AutoRouter(&controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test_:id([0-9]+)", &controllers.MainController{}, "*:Test")

	beego.Router("/subject", &controllers.SubjectController{})

	beego.Include(&controllers.ThreadController{})
	beego.Include(&controllers.Thread_logController{})

}
