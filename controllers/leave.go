package controllers

import (
	"todo/consts"
	"todo/models"
	"fmt"
)

type LeaveController struct {
	BaseController
}

// @router /leave/add [get,post]
func (c *LeaveController) Add() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		leave := new(models.Leave)
		c.Data["LeaveTypes"] = leave.GetLeaves()
		c.TplName = consts.LEAVE_ADD_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		leave := models.Leave{}
		err := c.ParseForm(&leave)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		user := c.GetSessionUser()
		leave.User = &user
		err = leave.Add()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("请假信息保存成功")
	}
}

// @router /leave/delete [post]
func (c *LeaveController) Delete() {
	LeaveId, _ := c.GetInt("LeaveId")
	leave := new(models.Leave)
	err := leave.Delete(LeaveId)
	if err == nil {
		c.AjaxOk("请假信息删除成功")
	} else {
		c.AjaxErr(fmt.Sprintf("%s",err))
	}
}

// @router /leave/update [get,post]
func (c *LeaveController) Update() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		LeaveId, err := c.GetInt("LeaveId")
		if err == nil && LeaveId > 0 {
			leave := new(models.Leave)
			one := leave.One(LeaveId)
			c.Data["Leave"] = one
			c.Data["LeaveTypes"] = leave.GetLeaves()
		}
		c.TplName = consts.LEAVE_UPDATE_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		leave := models.Leave{}
		err := c.ParseForm(&leave)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		user := c.GetSessionUser()
		leave.User = &user
		err = leave.Update()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("请假信息更新成功")
	}
}

// @router /leave/list [get]
func (c *LeaveController) List() {
	leave := new(models.Leave)
	c.Data["LeaveTypes"] = leave.GetLeaves()
	c.TplName = consts.LEAVE_INDEX_TPLName
}

// @router /leave/doList [get]
func (c *LeaveController) DoList() {
	page := c.GetPage()
	leaveType := c.GetString("leaveType")
	leaveStartTime := c.GetString("leaveStartTime")
	leaveEndTime := c.GetString("leaveEndTime")
	userId := c.GetSessionUser().UserId
	leave := new(models.Leave)
	resultPage := leave.List(page, leaveType, leaveStartTime, leaveEndTime, userId)
	c.Page(resultPage)
}
