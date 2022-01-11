package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego_study/controllers:ThreadController"] = append(beego.GlobalControllerRouter["beego_study/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/thread/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:ThreadController"] = append(beego.GlobalControllerRouter["beego_study/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/thread/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:ThreadController"] = append(beego.GlobalControllerRouter["beego_study/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/thread/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:ThreadController"] = append(beego.GlobalControllerRouter["beego_study/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/thread/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:ThreadController"] = append(beego.GlobalControllerRouter["beego_study/controllers:ThreadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/thread/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"] = append(beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/thread_log/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"] = append(beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/thread_log/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"] = append(beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/thread_log/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"] = append(beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/thread_log/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"] = append(beego.GlobalControllerRouter["beego_study/controllers:Thread_logController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/thread_log/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
