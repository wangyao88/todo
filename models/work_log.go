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
