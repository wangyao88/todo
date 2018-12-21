package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Leave struct {
	LeaveId         int       `orm:"pk;auto"`
	LeaveNote       string    `orm:"size(200)"`
	LeaveType       string    `orm:"size(10)"`
	LeaveCreateTime time.Time `orm:"auto_now_add;type(date)"`
	LeaveUpdateTime time.Time `orm:"auto_now;type(date)"`
	LeaveStartTime  time.Time `orm:"type(date)"`
	LeaveEndTime    time.Time `orm:"type(date)"`
	LeaveSum        float32   `orm:"digits(4);decimals(1)"`
	User            *User     `orm:"rel(fk)"`
}

func (leave *Leave) GetLeaves() []string {
	leaves := make([]string, 5)
	leaves[0] = "事假"
	leaves[1] = "病假"
	leaves[2] = "年假"
	leaves[3] = "丧假"
	leaves[4] = "婚假"
	return leaves
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Leave))
}

func (leave *Leave) Add() error {
	orm := orm.NewOrm()
	_, err := orm.Insert(leave)
	return err
}

func (leave *Leave) Update() error {
	orm := orm.NewOrm()
	_, err := orm.Update(leave)
	return err
}

func (leave *Leave) List(page *Page, leaveType string, leaveStartTime string, leaveEndTime string, userId int) *Page {
	orm := orm.NewOrm()
	leaveTable := new(Leave)
	querySeter := orm.QueryTable(&leaveTable)
	if leaveType != "-1" {
		querySeter = querySeter.Filter("LeaveType", leaveType)
	}
	if leaveStartTime != "" {
		querySeter = querySeter.Filter("LeaveStartTime__gte", leaveStartTime)
	}
	if leaveEndTime != "" {
		querySeter = querySeter.Filter("LeaveEndTime__lte", leaveEndTime)
	}
	querySeter = querySeter.Filter("User__UserId", userId)
	var resultPage = new(Page)
	total, _ := querySeter.Count()
	resultPage.SetTotal(int(total))
	var results []*Leave
	querySeter.Limit(page.pageSize, page.GetOffset()).All(&results)
	resultPage.SetRows(results)
	return resultPage
}

func (leave *Leave) Delete(LeaveId int) error {
	orm := orm.NewOrm()
	leaveTable := new(Leave)
	_, err := orm.QueryTable(&leaveTable).Filter("LeaveId", LeaveId).Delete()
	return err
}

func (leave *Leave) One(LeaveId int) *Leave {
	orm := orm.NewOrm()
	leaveTable := new(Leave)
	var leaveResult = new(Leave)
	orm.QueryTable(&leaveTable).Filter("LeaveId", LeaveId).One(leaveResult)
	return leaveResult
}
