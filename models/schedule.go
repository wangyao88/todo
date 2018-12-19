package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Schedule struct {
	ScheduleId         int       `orm:"pk;auto"`
	ScheduleContent    string    `orm:"size(500)"`
	ScheduleCreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	ScheduleStartTime  time.Time `orm:"type(datetime)"`
	ScheduleEndTime    time.Time `orm:"type(datetime)"`
	ScheduleUpdateTime time.Time `orm:"auto_now;type(datetime)"`
	User               *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Schedule))
}
