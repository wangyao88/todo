package controllers

import (
	"todo/consts"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	sessionUser := c.GetSessionUser()
	c.Data["UserName"] = sessionUser.UserName
	c.Data["UserId"] = sessionUser.UserId
	c.TplName = consts.MAIN_TPLName
}
