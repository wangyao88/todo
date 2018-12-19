package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	UserId     int       `orm:"pk;auto"`
	UserName   string    `orm:"unique"`
	Password   string    `orm:"size(100)"`
	Salt       string    `orm:"size(10)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(User))
}

func (u *User)FindByName(userName string) User {
	var user User
	o := orm.NewOrm()
	o.QueryTable("tb_user").Filter("UserName", userName).One(&user)
	return user
}
