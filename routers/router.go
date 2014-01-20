package routers

import (
	"github.com/astaxie/beego"
	"uoj/controllers"
)

func init() {
	//	beego.RESTRouter("/object", &controllers.ObjectController{})

	beego.Router("/user/login", &controllers.UserController{}, `post:Login`)
	beego.Router("/user/logout", &controllers.UserController{}, `post:Logout`)
	beego.Router("/user/join", &controllers.UserController{}, `post:Join`)

	beego.Router("/u", &controllers.UserController{}, `get:Getu`)
	beego.Router("/u/:uid", &controllers.UserController{}, `get:Getu`)

	// beego.Router("/m", &controllers.MsgController{})
	// beego.Router("/m/u/:uid", &controllers.MsgController{})
	// beego.Router("/m/tag/:t", &controllers.MsgController{})

	beego.Router("/p", &controllers.ProblemController{})
	beego.Router("/p/tag/:tag", &controllers.ProblemController{}, `get:Tag`)
	beego.Router("/p/submit/:pid", &controllers.ProblemController{})
	beego.Router("/p/:pid", &controllers.ProblemController{})

	// beego.Router("/t", &controllers.TaskController{})
	// beego.Router("/t/:tid", &controllers.TaskController{})
	// beego.Router("/t/:tid/solve", &controllers.TaskController{})

	// // Admin
	// beego.Router("/admin/add", &controllers.ProblemController{})
	// beego.Router("/admin/add", &controllers.ProblemController{})

}
