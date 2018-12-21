package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Todo struct {
	TodoId            int       `orm:"pk;auto"`
	TodoTitle         string    `orm:"size(200)"`
	TodoContent       string    `orm:"size(500)"`
	TodoCreateTime    time.Time `orm:"auto_now_add;type(date)"`
	TodoExcpteEndTime time.Time `orm:"type(date)"`
	TodoRealyEndTime  time.Time `orm:"null;type(date)"`
	TodoUpdateTime    time.Time `orm:"auto_now;type(date)"`
	User              *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Todo))
}

func (todo *Todo) Add() error {
	orm := orm.NewOrm()
	_, err := orm.Insert(todo)
	return err
}

func (todo *Todo) Update() error {
	orm := orm.NewOrm()
	_, err := orm.Update(todo)
	return err
}

func (todo *Todo) List(page *Page, todoTitle string, todoContent string, todoStartTime string, todoEndTime string, userId int) *Page {
	orm := orm.NewOrm()
	todoTable := new(Todo)
	querySeter := orm.QueryTable(&todoTable)
	if todoTitle != "" {
		querySeter = querySeter.Filter("TodoTitle__icontains", todoTitle)
	}
	if todoContent != "" {
		querySeter = querySeter.Filter("TodoContent__icontains", todoContent)
	}
	if todoStartTime != "" {
		querySeter = querySeter.Filter("TodoCreateTime__gte", todoStartTime)
	}
	if todoEndTime != "" {
		querySeter = querySeter.Filter("TodoCreateTime__lte", todoEndTime)
	}
	//time.Parse("01/02/2006", "02/08/2015")
	querySeter = querySeter.Filter("User__UserId", userId)
	var resultPage = new(Page)
	total, _ := querySeter.Count()
	resultPage.SetTotal(int(total))
	var results []*Todo
	querySeter.Limit(page.pageSize, page.GetOffset()).All(&results)
	resultPage.SetRows(results)
	return resultPage
}

func (todo *Todo) Delete(TodoId int) error {
	orm := orm.NewOrm()
	todoTable := new(Todo)
	_, err := orm.QueryTable(&todoTable).Filter("TodoId", TodoId).Delete()
	return err
}

func (todo *Todo) One(TodoId int) *Todo {
	orm := orm.NewOrm()
	todoTable := new(Todo)
	var todoResult = new(Todo)
	orm.QueryTable(&todoTable).Filter("TodoId", TodoId).One(todoResult)
	return todoResult
}
