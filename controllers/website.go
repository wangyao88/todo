package controllers

import (
	"todo/consts"
)

type WebsiteController struct {
	BaseController
}

// @router /website/list [get]
func (c *WebsiteController) List() {
	c.TplName = consts.WEBSITE_INDEX_TPLName
}
