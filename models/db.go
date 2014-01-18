package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/ojh?charset=utf8")

	orm.RegisterModel(new(Problem))
	orm.RegisterModel(new(User))

	o = orm.NewOrm()
	o.Using("default")
}
