package controllers

import (
	"todo/consts"
	"todo/models"
	"fmt"
)

type ProjectController struct {
	BaseController
}

// @router /project/add [get,post]
func (c *ProjectController) Add() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		setCompanyForSelect(c)
		c.TplName = consts.PROJECT_ADD_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		project := models.Project{}
		err := c.ParseForm(&project)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s", err))
			return
		}
		companyId, _ := c.GetInt("Company.CompanyId")
		company := new(models.Company)
		companyResult := company.One(companyId)
		project.Company = companyResult
		user := c.GetSessionUser()
		project.User = &user
		err = project.Add()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s", err))
			return
		}
		c.AjaxOk("项目信息保存成功")
	}
}

// @router /project/delete [post]
func (c *ProjectController) Delete() {
	ProjectId, _ := c.GetInt("ProjectId")
	project := new(models.Project)
	err := project.Delete(ProjectId)
	if err == nil {
		c.AjaxOk("项目信息删除成功")
	} else {
		c.AjaxErr(fmt.Sprintf("%s",err))
	}
}

// @router /project/update [get,post]
func (c *ProjectController) Update() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		ProjectId, err := c.GetInt("ProjectId")
		if err == nil && ProjectId > 0 {
			project := new(models.Project)
			one := project.One(ProjectId)
			c.Data["Project"] = one
		}
		setCompanyForSelect(c)
		c.TplName = consts.PROJECT_UPDATE_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		project := models.Project{}
		err := c.ParseForm(&project)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		companyId, _ := c.GetInt("Company.CompanyId")
		company := new(models.Company)
		companyResult := company.One(companyId)
		project.Company = companyResult
		user := c.GetSessionUser()
		project.User = &user
		err = project.Update()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("项目信息更新成功")
	}
}

// @router /project/list [get]
func (c *ProjectController) List() {
	c.TplName = consts.PROJECT_INDEX_TPLName
}

// @router /project/doList [get]
func (c *ProjectController) DoList() {
	page := c.GetPage()
	projectName := c.GetString("projectName")
	userId := c.GetSessionUser().UserId
	project := new(models.Project)
	resultPage := project.List(page, projectName, userId)
	c.Page(resultPage)
}

func setCompanyForSelect(c *ProjectController) {
	userId := c.GetSessionUser().UserId
	company := new(models.Company)
	result := company.SimpleList(userId)
	c.Data["Companys"] = result
}

