package controllers

import (
	"todo/consts"
	"todo/models"
	"fmt"
)

type WorkLogController struct {
	BaseController
}

// @router /workLog/add [get,post]
func (c *WorkLogController) Add() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		setProjectForSelect(c)
		c.TplName = consts.WORKLOG_ADD_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		workLog := models.WorkLog{}
		err := c.ParseForm(&workLog)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s", err))
			return
		}
		projectId, _ := c.GetInt("Project.ProjectId")
		project := new(models.Project)
		projectResult := project.One(projectId)
		workLog.Project = projectResult
		user := c.GetSessionUser()
		workLog.User = &user
		err = workLog.Add()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s", err))
			return
		}
		c.AjaxOk("工作日志保存成功")
	}
}

// @router /workLog/delete [post]
func (c *WorkLogController) Delete() {
	WorkLogId, _ := c.GetInt("WorkLogId")
	workLog := new(models.WorkLog)
	err := workLog.Delete(WorkLogId)
	if err == nil {
		c.AjaxOk("工作日志删除成功")
	} else {
		c.AjaxErr(fmt.Sprintf("%s",err))
	}
}

// @router /workLog/update [get,post]
func (c *WorkLogController) Update() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		WorkLogId, err := c.GetInt("WorkLogId")
		if err == nil && WorkLogId > 0 {
			workLog := new(models.WorkLog)
			one := workLog.One(WorkLogId)
			c.Data["WorkLog"] = one
		}
		setProjectForSelect(c)
		c.TplName = consts.WORKLOG_UPDATE_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		workLog := models.WorkLog{}
		err := c.ParseForm(&workLog)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		projectId, _ := c.GetInt("Project.ProjectId")
		project := new(models.Project)
		projectResult := project.One(projectId)
		workLog.Project = projectResult
		user := c.GetSessionUser()
		workLog.User = &user
		err = workLog.Update()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("工作日志更新成功")
	}
}

// @router /workLog/list [get]
func (c *WorkLogController) List() {
	setProjectForSelect(c)
	c.TplName = consts.WORKLOG_INDEX_TPLName
}

// @router /workLog/doList [get]
func (c *WorkLogController) DoList() {
	page := c.GetPage()
	workLogTitle := c.GetString("workLogTitle")
	workLogContent := c.GetString("workLogContent")
	projectId, _ := c.GetInt("projectId")
	userId := c.GetSessionUser().UserId
	workLog := new(models.WorkLog)
	resultPage := workLog.List(page, workLogTitle, workLogContent, projectId, userId)
	c.Page(resultPage)
}

func setProjectForSelect(c *WorkLogController) {
	userId := c.GetSessionUser().UserId
	project := new(models.Project)
	result := project.SimpleList(userId)
	c.Data["Projects"] = result
}

