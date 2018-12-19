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
	ProjectStartTime   time.Time `orm:"type(datetime)"`
	ProjectEndTime     time.Time `orm:"null;type(datetime)"`
	User               *User     `orm:"rel(fk)"`
	Company            *Company  `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Project))
}
