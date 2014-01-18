package controllers

import (
	"fmt"
	//	"encoding/json"
	"github.com/astaxie/beego"
	"uoj/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Finish() {
	this.ServeJson()
}

func (this *UserController) Getu() {
	uid := this.Ctx.Input.Params[":uid"]
	if uid != "" {
		ob, err := models.LookUserById(uid)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		start := 0
		if this.GetString("page") != "" {
			start = models.User_PAGE_SIZE * (models.Str2n(this.GetString("page")) - 1)
		}
		obs := models.LookUsers(start)
		this.Data["json"] = obs
	}
}

func (this *UserController) Login() {
	uid := this.GetString("userid")
	pass := this.GetString("password")
	var user models.User
	var err error
	var res models.Ops
	if user, err = models.GetUserById(uid); err != nil {
		res.Info = "No Such User."
	} else {
		if models.PwCheck(pass, user.Pwd) {
			this.SetSession("uid", uid)
			user.Pwd = ""
			res.OK = true
			res.Info = user
		} else {
			res.Info = "Password Wrong."
		}
	}
	this.Data["json"] = res
}

func (this *UserController) Logout() {
	if this.GetSession("uid") != nil {
		this.Data["json"] = fmt.Sprint(this.GetSession("uid")) + "Exit."
		this.DelSession("uid")
	} else {
		this.Data["json"] = "No User."
	}
}

func (this *UserController) Fuck() {
	this.Data["json"] = "Fuck?"
}
