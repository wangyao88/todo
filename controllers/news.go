package controllers

import (
	"todo/models"
)

type NewsController struct {
	BaseController
}

// @router /news/list [get]
func (c *NewsController) DoList() {
	var offset int
	offsetSession := c.GetSession("offset")
	if offsetSession == nil {
		offset = 0
	} else {
		offset = offsetSession.(int)
	}
	news := new(models.News)
	offset, newsResult := news.List(offset)
	c.SetSession("offset", offset)
	c.AjaxData(newsResult)
}
