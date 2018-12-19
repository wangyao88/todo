package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) AjaxOk(str string) {
	c.Data["json"] = ajax(str, 1)
	c.ServeJSON()
	c.StopRun()
}
func (c *BaseController) AjaxErr(str string) {
	c.Data["json"] = ajax(str, 0)
	c.ServeJSON()
	c.StopRun()
}

func ajax(str string, status int) (map[string]interface{}) {
	json := make(map[string]interface{})
	json["status"] = status
	json["msg"] = str
	return json
}
