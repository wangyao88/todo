package controllers

import (
	"time"
	"todo/consts"
	"todo/models"
)

type HomeController struct {
	BaseController
}

type date struct {
	Year  string
	Month string
	Day   string
}

// @router /home/index [get]
func (c *HomeController) Index() {
	setDate(c)
	setData(c)
	c.TplName = consts.HOME_INDEX_TPLName
}
func setData(c *HomeController) {
	todo := new(models.Todo)
	userId := c.GetSessionUser().UserId
	c.Data["Todos"] = todo.ListForTodo(userId)
	c.Data["UnTodos"] = todo.ListForUnTodo(userId)
}

func setDate(c *HomeController) {
	date := new(date)
	date.Year = time.Now().Format("2006")
	date.Month = time.Now().Format("01")
	date.Day = time.Now().Format("02")
	c.Data["Date"] = date
}
