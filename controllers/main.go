package controllers

import (
	"github.com/astaxie/beego"
	"todo/consts"
	"todo/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sessionUser := c.GetSession(consts.USER_KEY_IN_SESSION)
	user, ok := sessionUser.(models.User)
	if ok {
		fmt.Printf("%v",user)
		c.Data["UserName"] = user.UserName
		c.Data["UserId"] = user.UserId
	}
	c.TplName = consts.MAIN_TPLName
}
