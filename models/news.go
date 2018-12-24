package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type News struct {
	NewsId    string `orm:"pk;size(36)"`
	NewsTitle string `orm:"size(100)"`
	NewsUrl   string `orm:"size(500)"`
	NewsDate  string `orm:"type(date)"`
}

var total int

func init() {
	orm.RegisterModelWithPrefix("tb_", new(News))
}

func (company *News) List(offset int) (int, []*News) {
	orm := orm.NewOrm()
	newsTable := new(News)
	querySeter := orm.QueryTable(&newsTable)
	if offset >= total {
		offset = 0
	}
	var results []*News
	querySeter.Filter("NewsDate", getNowDate()).Limit(5,offset).All(&results)
	return offset+5, results
}

func getNewsTotal() int {
	orm := orm.NewOrm()
	newsTable := new(News)
	querySeter := orm.QueryTable(&newsTable)
	count, _ := querySeter.Filter("NewsDate", getNowDate()).Count()
	return int(count)
}

func (news *News)SetTotalTicker() {
	total = getNewsTotal()
	ticker := time.NewTicker(time.Hour)
	for _ = range ticker.C {
		total = getNewsTotal()
	}
}
