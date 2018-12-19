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
		return
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

// @router /company/delete [get,post]
func (c *CompanyController) Delete() {

}

// @router /company/update [get,post]
func (c *CompanyController) Update() {

}

// @router /company/one [get,post]
func (c *CompanyController) One() {

}

// @router /company/list [get]
func (c *CompanyController) List() {
	c.TplName = consts.COMPANY_INDEX_TPLName
}
