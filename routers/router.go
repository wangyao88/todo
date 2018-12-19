package routers

import (
	"todo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Include(&controllers.LoginController{})
}
