package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	//"strings"
)

var PRO_PAGE_SIZE int = 8

type Problem struct {
	Id    int    `orm:"pk;column(problem_id)"`
	Title string `orm:"index"`
	Ctx   string `orm:"type(text)"`
	Tag   string
	Tmlim int  `orm:"column(time_limit)"`
	Mnlim int  `orm:"column(memory_limit)"`
	Tot   int  `orm:"column(submit)"`
	Ac    int  `orm:"column(accepted)"`
	Live  bool `orm:"default(false")"`
}

func init() {
	//	Old2New()
}

func ReadPros(idx int) (pros []Problem) {
	qs := o.QueryTable("problem")
	qs.Filter("Id__gte", idx).Limit(PRO_PAGE_SIZE, idx).All(&pros, "Id", "Title", "Tag", "Tot", "Ac")
	return
}

func ReadProsByTag(tag string, idx int) (pros []Problem) {
	qs := o.QueryTable("problem")
	qs.Filter("Tag__icontains", tag).Limit(PRO_PAGE_SIZE, idx).OrderBy("Id").All(&pros, "Id", "Title", "Tag", "Tot", "Ac")
	return
}

func ReadProById(Pid int) (pro Problem, err error) {
	pro = Problem{Id: Pid}
	err = o.Read(&pro)
	if err != nil {
		return
	}
	return
}

func Old2New() {

	for pid := 1000; pid <= 1139; pid++ {
		fmt.Print("To Do ", pid, "...")
		var list orm.ParamsList
		// Problem.description -> Ctx
		num, err := o.Raw("SELECT description,input,output,sample_input,sample_output,hint FROM problem WHERE problem_id = ?", pid).ValuesFlat(&list)
		if err == nil && num > 0 {
			var ctx string
			ctx = fmt.Sprintln("##Description\n", fmt.Sprintln(list[0]),
				"\n##Input\n", fmt.Sprintln(list[1]),
				"\n##Output\n", fmt.Sprintln(list[2]),
				"\n##Sample Input\n<pre>", fmt.Sprintln(list[3]),
				"\n</pre>\n##Sample Output\n<pre>", fmt.Sprintln(list[4]), "</pre>")
			if len(fmt.Sprint(list[5])) > 3 {
				ctx = ctx + "\n##Hint\n" + fmt.Sprint(list[5])
			}
			//	ctx = strings.Replace(ctx, "\n", "\r\n", -1)
			_, err := o.Raw("UPDATE problem SET ctx = ? WHERE problem_id = ?", ctx, pid).Exec()
			if err == nil {
				fmt.Println("changed OK.")
			}
		}
		// Source -> Tag
		num, err = o.Raw("SELECT source FROM problem WHERE problem_id = ?", pid).ValuesFlat(&list)
		if err == nil && num > 0 {
			var ctx string
			ctx = fmt.Sprintln("[<", fmt.Sprintln(list[0]), ">]")
			_, err := o.Raw("UPDATE problem SET tag = ? WHERE problem_id = ?", ctx, pid).Exec()
			if err == nil {
				fmt.Println("changed OK.")
			}
		}
	}
}
