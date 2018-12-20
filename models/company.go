package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Company struct {
	CompanyId          int       `orm:"pk;auto"`
	CompanyName        string    `orm:"size(100)"`
	CompanyDescription string    `orm:"size(500)"`
	CompanyJob         string    `orm:"size(200)"`
	CompanyInTime      time.Time `orm:"type(date)"`
	CompanyOutTime     time.Time `orm:"null;type(date)"`
	User               *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Company))
}

func (company *Company) Add() error {
	orm := orm.NewOrm()
	_, err := orm.Insert(company)
	return err
}

func (company *Company) Update() error {
	orm := orm.NewOrm()
	_, err := orm.Update(company)
	return err
}

func (company *Company) List(page *Page, companyName string, userId int) *Page {
	orm := orm.NewOrm()
	companyTable := new(Company)
	querySeter := orm.QueryTable(&companyTable)
	if companyName != "" {
		querySeter = querySeter.Filter("CompanyName__icontains", companyName)
	}
	querySeter = querySeter.Filter("User__UserId", userId)
	var resultPage = new(Page)
	total, _ := querySeter.Count()
	resultPage.SetTotal(int(total))
	var results []*Company
	querySeter.Limit(page.pageSize, page.GetOffset()).All(&results)
	resultPage.SetRows(results)
	return resultPage
}

func (company *Company) Delete(CompanyId int) error {
	orm := orm.NewOrm()
	companyTable := new(Company)
	_, err := orm.QueryTable(&companyTable).Filter("CompanyId", CompanyId).Delete()
	return err
}

func (company *Company) One(CompanyId int) *Company {
	orm := orm.NewOrm()
	companyTable := new(Company)
	var companyResult = new(Company)
	orm.QueryTable(&companyTable).Filter("CompanyId", CompanyId).One(companyResult)
	return companyResult
}

func (company *Company) SimpleList(userId int) []*Company {
	orm := orm.NewOrm()
	companyTable := new(Company)
	var results []*Company
	orm.QueryTable(&companyTable).Filter("User__UserId", userId).All(&results, "CompanyId" ,"CompanyName")
	return results
}
