package routers

import (
	"beego_study/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 路由配置 https://www.jianshu.com/p/417fcc5841f8
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/test", &controllers.MainController{}, "*:Test")

	beego.Include(&controllers.ThreadController{})

}
