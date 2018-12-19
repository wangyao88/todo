package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"todo/consts"
	_ "todo/routers"
)

func init() {
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbUser := beego.AppConfig.String("dbUser")
	dbPassword := beego.AppConfig.String("dbPassword")
	dbName := beego.AppConfig.String("dbName")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDriver(consts.MYSQL_DATABASE_DRIVERNAME, orm.DRMySQL)
	orm.RegisterDataBase(consts.DEFAULT_DATABASE_ALIASNAME, consts.MYSQL_DATABASE_DRIVERNAME, dsn)
}

var FilterLogin = func(ctx *context.Context) {
	user := ctx.Input.Session(consts.USER_KEY_IN_SESSION)
	if user == nil && ctx.Request.RequestURI != consts.LOGIN_URL {
		ctx.Redirect(consts.REDIRECT_STATUS_CODE, consts.LOGIN_URL)
	}
}

func autoCreateTable() {
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	autoCreateTable()
	orm.RunCommand()
	orm.Debug = true
	beego.InsertFilter(consts.LOGIN_FILTER_URL, beego.BeforeRouter, FilterLogin)
	beego.Run()
}
