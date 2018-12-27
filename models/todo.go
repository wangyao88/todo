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

func (todo *Todo) List(page *Page, todoTitle string, todoContent string, todoStartTime string, todoEndTime string, listType string, userId int) *Page {
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
	if listType != "" {
		if listType == "todo" {
			querySeter = querySeter.Filter("TodoRealyEndTime__isnull", true)
		}
		if listType == "unTodo" {
			querySeter = querySeter.Filter("TodoRealyEndTime__isnull", false).Filter("TodoRealyEndTime",getNowDate())
		}
	}
	//time.Parse("01/02/2006", "02/08/2015")
	querySeter = querySeter.Filter("User__UserId", userId)
	var resultPage = new(Page)
	total, _ := querySeter.Count()
	resultPage.SetTotal(int(total))
	var results []*Todo
	querySeter.OrderBy("-TodoCreateTime").Limit(page.pageSize, page.GetOffset()).All(&results)
	resultPage.SetRows(results)
	return resultPage
}

func (todo *Todo) ListForTodo(userId int) []*Todo {
	orm := orm.NewOrm()
	todoTable := new(Todo)
	querySeter := orm.QueryTable(&todoTable)
	querySeter = querySeter.Filter("TodoRealyEndTime__isnull", true).Filter("User__UserId", userId)
	var results []*Todo
	querySeter.Limit(5).All(&results,"TodoId","TodoTitle","TodoContent")
	return results
}

func (todo *Todo) ListForUnTodo(userId int) []*Todo {
	orm := orm.NewOrm()
	todoTable := new(Todo)
	querySeter := orm.QueryTable(&todoTable)
	querySeter = querySeter.Filter("TodoRealyEndTime__isnull", false).Filter("TodoRealyEndTime",getNowDate()).Filter("User__UserId", userId)
	var results []*Todo
	querySeter.Limit(5).All(&results,"TodoId","TodoTitle","TodoContent")
	return results
}

func getNowDate() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	return year + "-" + month + "-" + day;
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
