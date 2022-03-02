package routers

import (
	"beego_study/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 路由配置 https://www.jianshu.com/p/417fcc5841f8
	//beego.AutoRouter(&controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{}, "*:Register;post:RegisterPost")
	beego.Router("/register_:id([0-9]+)", &controllers.MainController{}, "*:RegisterOne") // 查询一个
	beego.Router("/register_list", &controllers.MainController{}, "*:RegisterList")
	beego.Router("/register_update", &controllers.MainController{}, "*:RegisterUpdate")
	beego.Router("/register_del", &controllers.MainController{}, "*:RegisterDel")
	beego.Router("/test_:id([0-9]+)", &controllers.MainController{}, "*:Test")

	beego.Router("/login", &controllers.MainController{}, "*:Login;post:LoginPost")
	beego.Router("/subject", &controllers.SubjectController{})

	beego.Include(&controllers.ThreadController{})
	beego.Include(&controllers.Thread_logController{})

}
