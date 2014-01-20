package controllers

import (
	"fmt"
	"time"
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

func (this *UserController) Join() {
	uid, team, nick, pwd, pwd2, email, sid, sname :=
		this.GetString("userid"), this.GetString("team"),
		this.GetString("nickname"), this.GetString("password"),
		this.GetString("password2"), this.GetString("emailaddr"),
		this.GetString("studentid"), this.GetString("realname")
	var res models.Ops
	switch {
	case pwd != pwd2 || models.BadMail(email):
		res.Info = 202 // 信息不匹配
	case models.HasNull(uid, pwd, email, sid, sname):
		res.Info = 201 // 信息不完整
	default:
		u, _ := models.GetUserById(uid)
		if len(u.Pwd) != 0 {
			res.Info = 210 // 用户已存在
		} else {
			u := models.User{
				uid, nick, pwd, email, time.Now(), team,
				sid, sname, 0, 0, 0, fmt.Sprintf("%s", this.Ctx.Request.RemoteAddr), true,
			}
			if err := models.AddUser(u); err == nil {
				res.OK = true
				res.Info = u
			} else {
				res.Info = err
			}
		}
	}
	this.Data["json"] = res
}

func (this *UserController) Login() {
	uid := this.GetString("userid")
	pass := this.GetString("password")
	var user models.User
	var err error
	var res models.Ops
	if user, err = models.GetUserById(uid); err != nil {
		res.Info = 200
	} else {
		if models.PwCheck(pass, user.Pwd) {
			this.SetSession("uid", uid)
			user.Pwd = ""
			res.OK = true
			res.Info = user
		} else {
			res.Info = 202
		}
	}
	this.Data["json"] = res
}

func (this *UserController) Logout() {
	if this.GetSession("uid") != nil {
		this.Data["json"] = fmt.Sprint(this.GetSession("uid")) + "Exit."
		this.DelSession("uid")
	} else {
		this.Data["json"] = 200
	}
}

func (this *UserController) Fuck() {
	this.Data["json"] = "Fuck?"
}
