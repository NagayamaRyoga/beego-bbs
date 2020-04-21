package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqlConn := beego.AppConfig.String("mysqlConn")
	if err := orm.RegisterDataBase("default", "mysql", mysqlConn); err != nil {
		panic(err)
	}
}
