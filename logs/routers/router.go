package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"logs/controllers/admin"
	"logs/controllers/client"
	"logs/controllers/item"
	"logs/controllers/logs"
)

func init() {
	// 验证用户是否已经登录
	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)
	beego.Router("/", &admin.UserController{}, "*:Login")
	beego.Router("/admin/test", &admin.UserController{}, "*:Test")
	beego.Router("/admin/console", &admin.UserController{}, "*:Console")
	beego.Router("/admin/index", &admin.UserController{}, "*:Index")
	beego.Router("/admin/login", &admin.UserController{}, "*:Login")
	beego.Router("/admin/exit", &admin.UserController{}, "*:Exit")
	beego.Router("/client/index", &client.ClientController{}, "*:Index")
	beego.Router("/item/index", &item.ItemController{}, "*:Index")
	beego.Router("/logs/index", &logs.LogsController{}, "*:Index")
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("username").(string)

	if !ok && ctx.Request.RequestURI != "/admin/login" {
		ctx.Redirect(302, "/admin/login")
	}
}
