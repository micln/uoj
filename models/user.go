package models

import (
	//	"fmt"
	"github.com/Unknwon/com"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Uid     string `orm:"pk;column(user_id)"`
	Nick    string
	Pwd     string `orm:"column(password)"`
	Email   string
	RegTime time.Time `orm:"column(accesstime)"`
	Team    string
	Sid     string `orm:"column(ssid)"`
	Sname   string `orm:"column(ssname)"`
	Submit  int
	Ac      int `orm:"column(solved)"`
	Power   int
	Ip      string
	Live    bool
}

// Sample field ,Look in Other
type Userspl struct {
	Uid    string `orm:"pk;column(user_id)"`
	Nick   string
	Email  string
	Team   string
	Submit int
	Ac     int `orm:"column(solved)"`
}

const (
	User_PAGE_SIZE = 10
)

// Get all field
func GetUserById(uid string) (user User, err error) {
	// qs := o.QueryTable("user")
	// qs.Filter("Uid", uid).One(&user)
	user = User{Uid: uid}
	err = o.Read(&user)
	return
}

// Sample field ,Look in Other
func LookUserById(uid string) (user Userspl, err error) {
	usera := User{}
	qs := o.QueryTable("user")
	qs.Filter("Uid", uid).One(&usera, "Uid", "Nick", "Email", "Team", "Submit", "Ac")
	user = Userspl{usera.Uid, usera.Nick, usera.Email, usera.Team, usera.Submit, usera.Ac}
	return
}

func GetUsers(idx int) (users []User) {
	qs := o.QueryTable("user")
	qs.Limit(User_PAGE_SIZE, idx).OrderBy("-Ac").All(&users)
	return
}

func LookUsers(idx int) []Userspl {
	usersa := []User{}
	qs := o.QueryTable("user")
	qs.Limit(User_PAGE_SIZE, idx).OrderBy("-Ac").All(&usersa)
	users := [User_PAGE_SIZE]Userspl{}
	for k, v := range usersa {
		users[k] = Userspl{v.Uid, v.Nick, v.Email, v.Team, v.Submit, v.Ac}
	}
	return users[:]
}

func PwGen(pass string) string {
	salt := "8bzm"
	hash := com.Base64Encode(Sha120(com.Md5(pass)+salt) + salt)
	return hash
}

func PwCheck(pass, saved string) bool {
	svd, err := com.Base64Decode(saved)
	if err != nil {
		return false
	}
	salt := com.SubStr(svd, -3, 8)
	hash := com.Base64Encode(Sha120(com.Md5(pass)+salt) + salt)
	if hash == saved {
		return true
	}
	return false
}

func Test() interface{} {
	var list []orm.ParamsList
	num, err := o.Raw("SELECT * FROM solution WHERE user_id = ?", "zhr").ValuesList(&list)
	if err == nil && num > 0 {
		return list
	}
	return nil
}
