package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Project struct {
	ProjectId          int       `orm:"pk;auto"`
	ProjectName        string    `orm:"size(100)"`
	ProjectDescription string    `orm:"size(500)"`
	ProjectJob         string    `orm:"size(200)"`
	ProjectStartTime   time.Time `orm:"type(date)"`
	ProjectEndTime     time.Time `orm:"null;type(date)"`
	User               *User     `orm:"rel(fk)"`
	Company            *Company  `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Project))
}

func (project *Project) Add() error {
	orm := orm.NewOrm()
	_, err := orm.Insert(project)
	return err
}

func (project *Project) Update() error {
	orm := orm.NewOrm()
	_, err := orm.Update(project)
	return err
}

func (project *Project) List(page *Page, projectName string, userId int) *Page {
	orm := orm.NewOrm()
	projectTable := new(Project)
	querySeter := orm.QueryTable(&projectTable)
	if projectName != "" {
		querySeter = querySeter.Filter("ProjectName__icontains", projectName)
	}
	querySeter = querySeter.Filter("User__UserId", userId)
	var resultPage = new(Page)
	total, _ := querySeter.Count()
	resultPage.SetTotal(int(total))
	var results []*Project
	querySeter.Limit(page.pageSize, page.GetOffset()).RelatedSel().All(&results)
	resultPage.SetRows(results)
	return resultPage
}

func (project *Project) Delete(ProjectId int) error {
	orm := orm.NewOrm()
	projectTable := new(Project)
	_, err := orm.QueryTable(&projectTable).Filter("ProjectId", ProjectId).Delete()
	return err
}

func (project *Project) One(ProjectId int) *Project {
	orm := orm.NewOrm()
	projectTable := new(Project)
	var projectResult = new(Project)
	orm.QueryTable(&projectTable).Filter("ProjectId", ProjectId).One(projectResult)
	return projectResult
}

func (project *Project) SimpleList(userId int) []*Project {
	orm := orm.NewOrm()
	projectTable := new(Project)
	var results []*Project
	orm.QueryTable(&projectTable).Filter("User__UserId", userId).All(&results, "ProjectId" ,"ProjectName")
	return results
}