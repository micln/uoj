package models

import (
	//	"fmt"
	"github.com/Unknwon/com"
	//	"github.com/astaxie/beego"
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

// Sample field ,looked in the third view
type Usertpl struct {
	UserId      string
	Nickname    string
	EmailAddr   string
	Team        string
	Submit      int
	Solved      int
	StudentId   string
	StudentName string
}

const (
	User_PAGE_SIZE = 10
)

// Add a user
func AddUser(u User) (err error) {
	_, err = o.Insert(u)
	return
}

// Get all fields for someone
func GetUserById(uid string) (user User, err error) {
	user = User{Uid: uid}
	err = o.Read(&user)
	// o.QueryTable("user").Filter("Uid", uid).One(&user)
	return
}

// Get all fields for all users
func GetUsers(idx int) (users []User) {
	o.QueryTable("user").Limit(User_PAGE_SIZE, idx).OrderBy("-Ac").All(&users)
	return
}

// Get some field for someone in the third view
func LookUserById(uid string) (user Usertpl, err error) {
	usera := User{}
	o.QueryTable("user").Filter("Uid", uid).One(&usera, "Uid", "Nick", "Email", "Team", "Submit", "Ac", "Sid", "Sname")
	user = Usertpl{usera.Uid, usera.Nick, usera.Email, usera.Team, usera.Submit, usera.Ac, usera.Sid, usera.Sname}
	return
}

// Get some field for all in the third view
func LookUsers(idx int) []Usertpl {
	usersa, users := []User{}, [User_PAGE_SIZE]Usertpl{}
	o.QueryTable("user").Limit(User_PAGE_SIZE, idx).OrderBy("-Ac").All(&usersa)
	for k, v := range usersa {
		users[k] = Usertpl{v.Uid, v.Nick, v.Email, v.Team, v.Submit, v.Ac, v.Sid, v.Sname}
	}
	return users[:]
}

// Look users by studentId Prefix
func LookUsersPx(uid string) []Usertpl {
	usera, users := []User{}, []Usertpl{}
	o.QueryTable("user").Filter("Sid__icontains", uid).OrderBy("-Ac").All(&usera)
	for _, v := range usera {
		users = append(users, Usertpl{v.Uid, v.Nick, v.Email, v.Team, v.Submit, v.Ac, v.Sid, v.Sname})
	}
	return users
}

// encode the password
func PwGen(pass string) (hash string) {
	salt := "8bzm"
	hash = com.Base64Encode(Sha120(com.Md5(pass)+salt) + salt)
	return
}

// check whether the password is correct
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
