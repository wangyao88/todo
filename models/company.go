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
	CompanyInTime      time.Time `orm:"type(datetime)"`
	CompanyOutTime     time.Time `orm:"null;type(datetime)"`
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