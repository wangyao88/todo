package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["todo/controllers:CompanyController"] = append(beego.GlobalControllerRouter["todo/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/company/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:CompanyController"] = append(beego.GlobalControllerRouter["todo/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/company/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:CompanyController"] = append(beego.GlobalControllerRouter["todo/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "DoList",
            Router: `/company/doList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:CompanyController"] = append(beego.GlobalControllerRouter["todo/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/company/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:CompanyController"] = append(beego.GlobalControllerRouter["todo/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/company/update`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:LoginController"] = append(beego.GlobalControllerRouter["todo/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:LoginController"] = append(beego.GlobalControllerRouter["todo/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:ProjectController"] = append(beego.GlobalControllerRouter["todo/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/project/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:ProjectController"] = append(beego.GlobalControllerRouter["todo/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/project/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:ProjectController"] = append(beego.GlobalControllerRouter["todo/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DoList",
            Router: `/project/doList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:ProjectController"] = append(beego.GlobalControllerRouter["todo/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/project/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:ProjectController"] = append(beego.GlobalControllerRouter["todo/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/project/update`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:WorkLogController"] = append(beego.GlobalControllerRouter["todo/controllers:WorkLogController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/workLog/add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:WorkLogController"] = append(beego.GlobalControllerRouter["todo/controllers:WorkLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/workLog/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:WorkLogController"] = append(beego.GlobalControllerRouter["todo/controllers:WorkLogController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: `/workLog/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:WorkLogController"] = append(beego.GlobalControllerRouter["todo/controllers:WorkLogController"],
        beego.ControllerComments{
            Method: "DoList",
            Router: `/workLog/doList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:WorkLogController"] = append(beego.GlobalControllerRouter["todo/controllers:WorkLogController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/workLog/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["todo/controllers:WorkLogController"] = append(beego.GlobalControllerRouter["todo/controllers:WorkLogController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/workLog/update`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
