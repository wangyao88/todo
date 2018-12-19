package controllers

import "todo/consts"

type CompanyController struct {
	BaseController
}

// @router /company/add [get,post]
func (c *CompanyController) Add() {
	if c.Ctx.Request.Method == consts.REQUEST_GET_METHOD {
		c.TplName = consts.COMPANY_ADD_TPLName
	} else {

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
