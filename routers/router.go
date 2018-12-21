package routers

import (
	"github.com/astaxie/beego"
	"todo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.CompanyController{})
	beego.Include(&controllers.ProjectController{})
	beego.Include(&controllers.WorkLogController{})
}
