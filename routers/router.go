package routers

import (
	"beego_study/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 路由配置 https://www.jianshu.com/p/417fcc5841f8
	// beego.AutoRouter(&controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{}, "*:Index")
	beego.Router("/test_:id([0-9]+)", &controllers.MainController{}, "*:Test")

	beego.Router("/subject", &controllers.SubjectController{})

	// 用户注册
	beego.Router("/register", &controllers.RegisterController{})
	// 查询或更改一个注册用户
	beego.Router("/register_:id([0-9]+)", &controllers.RegisterController{}, "*:RegisterOne;post:RegisterUpdate")
	// 查询所有的注册用户
	beego.Router("/register_list", &controllers.RegisterController{}, "*:RegisterList")
	// 删除一个注册用户
	beego.Router("/register_delete_:id([0-9]+)", &controllers.RegisterController{}, "*:RegisterDelete")

	// 用户登录
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article_list", &controllers.ArticleController{}, "*:ArticleList")
	beego.Router("/article_update", &controllers.ArticleController{}, "*:ArticleUpdate")
	beego.Router("/article_:id([0-9]+)", &controllers.ArticleController{}, "*:ArticleInfo")
	beego.Router("/article_delete_:id([0-9]+)", &controllers.ArticleController{}, "*:ArticleDelete")

	beego.Include(&controllers.ThreadController{})
	beego.Include(&controllers.Thread_logController{})
}
