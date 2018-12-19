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
            AllowHTTPMethods: []string{"get","post"},
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
            Method: "One",
            Router: `/company/one`,
            AllowHTTPMethods: []string{"get","post"},
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

}
