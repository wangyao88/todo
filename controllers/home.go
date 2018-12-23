package controllers

import (
	"todo/consts"
)

type HomeController struct {
	BaseController
}

// @router /home/index [get]
func (c *HomeController) Index() {
	c.TplName = consts.HOME_INDEX_TPLName
}
