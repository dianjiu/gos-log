package routers

import (
	"github.com/astaxie/beego"
	"logs/controllers"
	"logs/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/test", &admin.UserController{}, "*:Test")
	beego.Router("/admin/login/?:user_name", &admin.UserController{}, "*:Login")
	beego.Router("/admin/login/:username:/:password:", &admin.UserController{}, "*:Login")
}
