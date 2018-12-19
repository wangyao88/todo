package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Todo struct {
	TodoId            int       `orm:"pk;auto"`
	TodoTitle         string    `orm:"size(200)"`
	TodoContent       string    `orm:"size(500)"`
	TodoCreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	TodoExcpteEndTime time.Time `orm:"type(datetime)"`
	TodoRealyEndTime  time.Time `orm:"null;type(datetime)"`
	TodoUpdateTime    time.Time `orm:"auto_now;type(datetime)"`
	User              *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Todo))
}
