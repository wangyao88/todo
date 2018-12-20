package controllers

import (
	"todo/consts"
	"todo/models"
	"fmt"
)

type CompanyController struct {
	BaseController
}

// @router /company/add [get,post]
func (c *CompanyController) Add() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		c.TplName = consts.COMPANY_ADD_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		company := models.Company{}
		err := c.ParseForm(&company)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		user := c.GetSessionUser()
		company.User = &user
		err = company.Add()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("公司信息保存成功")
	}
}

// @router /company/delete [post]
func (c *CompanyController) Delete() {
	CompanyId, _ := c.GetInt("CompanyId")
	company := new(models.Company)
	err := company.Delete(CompanyId)
	if err == nil {
		c.AjaxOk("公司信息删除成功")
	} else {
		c.AjaxErr(fmt.Sprintf("%s",err))
	}
}

// @router /company/update [get,post]
func (c *CompanyController) Update() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		CompanyId, err := c.GetInt("CompanyId")
		if err == nil && CompanyId > 0 {
			company := new(models.Company)
			one := company.One(CompanyId)
			c.Data["Company"] = one
		}
		c.TplName = consts.COMPANY_UPDATE_TPLName
	}
	if c.Ctx.Request.Method == consts.REQUEST_POST_METHOD {
		company := models.Company{}
		err := c.ParseForm(&company)
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		user := c.GetSessionUser()
		company.User = &user
		err = company.Update()
		if err != nil {
			c.AjaxErr(fmt.Sprintf("%s",err))
			return
		}
		c.AjaxOk("公司信息更新成功")
	}
}

// @router /company/list [get]
func (c *CompanyController) List() {
	c.TplName = consts.COMPANY_INDEX_TPLName
}

// @router /company/doList [get]
func (c *CompanyController) DoList() {
	page := c.GetPage()
	companyName := c.GetString("companyName")
	userId := c.GetSessionUser().UserId
	company := new(models.Company)
	resultPage := company.List(page, companyName, userId)
	c.Page(resultPage)
}
