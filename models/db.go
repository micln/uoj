package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// global ormer
var o orm.Ormer

// DB_MYSQL start.
func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:passwd@/ojh?charset=utf8")

	orm.RegisterModel(new(Problem), new(User))

	o = orm.NewOrm()
	o.Using("default")
}
