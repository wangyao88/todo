package controllers

import (
	"todo/consts"
	"todo/models"
	"todo/utils"
)

type LoginController struct {
	BaseController
}

// @router /login [get,post]
func (c *LoginController) Login() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		c.TplName = consts.LOGIN_TPLName
		//var user models.User
		//var encrypt utils.Encrypt
		//user.Registe("wy",encrypt.Str2md5("123"+consts.DEFAULT_SALT))
	} else {
		userName := c.GetString("UserName")
		password := c.GetString("Password")
		var user models.User
		userPersist := user.FindByName(userName)
		var encrypt utils.Encrypt
		if userPersist.Password == encrypt.Str2md5(password+userPersist.Salt) {
			c.SetSession(consts.USER_KEY_IN_SESSION, userPersist)
			c.AjaxOk("登陆成功")
		} else {
			c.AjaxErr("用户名或密码不正确")
		}
	}
}

// @router /logout [post]
func (c *LoginController) Logout() {
	c.DelSession(consts.USER_KEY_IN_SESSION)
	c.AjaxOk("注销成功")
}
