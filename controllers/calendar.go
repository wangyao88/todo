package controllers

import (
	"todo/consts"
)

type CalendarController struct {
	BaseController
}

// @router /calendar/list [get]
func (c *CalendarController) List() {
	c.TplName = consts.CALENDAR_INDEX_TPLName
}
