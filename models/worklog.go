package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type WorkLog struct {
	WorkLogId         int       `orm:"pk;auto"`
	WorkLogTitle      string    `orm:"size(200)"`
	WorkLogContent    string    `orm:"size(2000)"`
	WorkLogCreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	WorkLogUpdateTime time.Time `orm:"auto_now;type(datetime)"`
	User              *User     `orm:"rel(fk)"`
	Project           *Project  `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(WorkLog))
}

func (workLog *WorkLog) Add() error {
	orm := orm.NewOrm()
	_, err := orm.Insert(workLog)
	return err
}

func (workLog *WorkLog) Update() error {
	orm := orm.NewOrm()
	_, err := orm.Update(workLog)
	return err
}

func (workLog *WorkLog) List(page *Page, workLogTitle string, workLogContent string, projectId int, userId int) *Page {
	orm := orm.NewOrm()
	workLogTable := new(WorkLog)
	querySeter := orm.QueryTable(&workLogTable)
	if workLogTitle != "" {
		querySeter = querySeter.Filter("WorkLogTitle__icontains", workLogTitle)
	}
	if workLogContent != "" {
		querySeter = querySeter.Filter("WorkLogContent__icontains", workLogContent)
	}
	if projectId != -1 {
		querySeter = querySeter.Filter("Project__ProjectId", projectId)
	}
	querySeter = querySeter.Filter("User__UserId", userId)
	var resultPage = new(Page)
	total, _ := querySeter.Count()
	resultPage.SetTotal(int(total))
	var results []*WorkLog
	querySeter.Limit(page.pageSize, page.GetOffset()).RelatedSel().All(&results)
	resultPage.SetRows(results)
	return resultPage
}

func (workLog *WorkLog) Delete(WorkLogId int) error {
	orm := orm.NewOrm()
	workLogTable := new(WorkLog)
	_, err := orm.QueryTable(&workLogTable).Filter("WorkLogId", WorkLogId).Delete()
	return err
}

func (workLog *WorkLog) One(WorkLogId int) *WorkLog {
	orm := orm.NewOrm()
	workLogTable := new(WorkLog)
	var workLogResult = new(WorkLog)
	orm.QueryTable(&workLogTable).Filter("WorkLogId", WorkLogId).One(workLogResult)
	return workLogResult
}
