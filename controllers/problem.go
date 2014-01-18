package controllers

import (
	//	"encoding/json"
	"github.com/astaxie/beego"
	"uoj/models"
)

type ProblemController struct {
	beego.Controller
}

func (this *ProblemController) Get() {
	Pid := this.Ctx.Input.Params[":pid"]
	if Pid != "" {
		ob, err := models.ReadProById(models.Str2n(Pid))
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		pid := 0
		if this.GetString("page") != "" {
			pid = models.PRO_PAGE_SIZE * (models.Str2n(this.GetString("page")) - 1)
		}
		obs := models.ReadPros(pid)
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *ProblemController) Tag() {
	tag := this.Ctx.Input.Params[":tag"]
	if tag != "" {
		tid := 0
		if this.GetString("page") != "" {
			tid = models.PRO_PAGE_SIZE * (models.Str2n(this.GetString("page")) - 1)
		}
		ob := models.ReadProsByTag(tag, tid)
		this.Data["json"] = ob
	} else {
		obs := models.ReadPros(1010)
		this.Data["json"] = obs
	}
	this.ServeJson()
}
