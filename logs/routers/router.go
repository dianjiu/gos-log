package routers

import (
	"github.com/astaxie/beego"
	"logs/controllers"
	"logs/controllers/admin"
	"logs/controllers/client"
	"logs/controllers/item"
	"logs/controllers/logs"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/test", &admin.UserController{}, "*:Test")
	beego.Router("/admin/console", &admin.UserController{}, "*:Console")
	beego.Router("/admin/toLogin", &admin.UserController{}, "*:ToLogin")
	beego.Router("/admin/login/?:user_name", &admin.UserController{}, "*:Login")
	beego.Router("/admin/login/:username:/:password:", &admin.UserController{}, "*:Login")
	beego.Router("/client/index", &client.ClientController{}, "*:Index")
	beego.Router("/item/index", &item.ItemController{}, "*:Index")
	beego.Router("/logs/index", &logs.LogsController{}, "*:Index")
}
