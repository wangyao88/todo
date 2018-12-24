package controllers

import (
	"todo/consts"
	"todo/models"
	"fmt"
)

type TodoController struct {
	BaseController
}

// @router /todo/add [get,post]
func (c *TodoController) Add() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		c.TplName = consts.TODO_ADD_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		todo := models.Todo{}
		err := c.ParseForm(&todo)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		user := c.GetSessionUser()
		todo.User = &user
		err = todo.Add()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("待办事项保存成功")
	}
}

// @router /todo/delete [post]
func (c *TodoController) Delete() {
	TodoId, _ := c.GetInt("TodoId")
	todo := new(models.Todo)
	err := todo.Delete(TodoId)
	if err == nil {
		c.AjaxOk("待办事项删除成功")
	} else {
		c.AjaxErr(fmt.Sprintf("%s",err))
	}
}

// @router /todo/update [get,post]
func (c *TodoController) Update() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		TodoId, err := c.GetInt("TodoId")
		if err == nil && TodoId > 0 {
			todo := new(models.Todo)
			one := todo.One(TodoId)
			c.Data["Todo"] = one
		}
		c.TplName = consts.TODO_UPDATE_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		todo := models.Todo{}
		err := c.ParseForm(&todo)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		user := c.GetSessionUser()
		todo.User = &user
		err = todo.Update()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("待办事项更新成功")
	}
}

// @router /todo/list [get]
func (c *TodoController) List() {
	c.TplName = consts.TODO_INDEX_TPLName
}

// @router /todo/doList [get]
func (c *TodoController) DoList() {
	page := c.GetPage()
	todoTitle := c.GetString("todoTitle")
	todoContent := c.GetString("todoContent")
	todoStartTime := c.GetString("todoStartTime")
	todoEndTime := c.GetString("todoEndTime")
	listType := c.GetString("listType")
	userId := c.GetSessionUser().UserId
	todo := new(models.Todo)
	resultPage := todo.List(page, todoTitle, todoContent, todoStartTime, todoEndTime,listType, userId)
	c.Page(resultPage)
}
