package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"todo/consts"
	"errors"
)

type User struct {
	UserId     int       `orm:"pk;auto"`
	UserName   string    `orm:"unique"`
	Password   string    `orm:"size(100)"`
	Salt       string    `orm:"size(10)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(User))
}

func (u *User) FindByName(userName string) User {
	var user User
	o := orm.NewOrm()
	o.QueryTable("tb_user").Filter("UserName", userName).One(&user)
	return user
}

func (u *User) Registe(userName string, password string) error {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("tb_user").Filter("UserName", userName).One(&user)
	if err != nil {
		if err == orm.ErrNoRows {
			registeUser := User{
				UserName: userName,
				Password: password,
				Salt: consts.DEFAULT_SALT,
			}
			_, err := o.Insert(&registeUser)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("用户已存在!")
	}
	return nil
}
