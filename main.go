package main

import (
	//	"fmt"
	//	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "uoj/models"
	_ "uoj/routers"
)

func init() {
	beego.SetLevel(beego.LevelTrace)
	orm.Debug = true
}

func main() {

	beego.Run()
}
