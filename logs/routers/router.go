package routers

import (
	"logs/controllers/admin"
	"logs/controllers/client"
	"logs/controllers/item"
	"logs/controllers/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	// 验证用户是否已经登录
	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)
	beego.Router("/", &admin.UserController{}, "*:Login")
	beego.Router("/admin/test", &admin.UserController{}, "*:Test")
	beego.Router("/admin/console", &admin.UserController{}, "*:Console")
	beego.Router("/admin/index", &admin.UserController{}, "*:Index")
	beego.Router("/admin/login", &admin.UserController{}, "*:Login")
	beego.Router("/admin/exit", &admin.UserController{}, "*:Exit")
	beego.Router("/client/index", &client.ClientController{}, "*:Index")
	beego.Router("/client/add", &client.ClientController{}, "*:Add")
	beego.Router("/client/delete", &client.ClientController{}, "*:Delete")
	beego.Router("/client/update", &client.ClientController{}, "*:Update")
	beego.Router("/client/queryAll", &client.ClientController{}, "*:QueryAll")
	beego.Router("/client/query", &client.ClientController{}, "*:Query")
	beego.Router("/client/queryPage", &client.ClientController{}, "*:QueryPage")
	beego.Router("/item/index", &item.ItemController{}, "*:Index")
	beego.Router("/logs/index", &logs.LogsController{}, "*:Index")
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("username").(string)

	if !ok && ctx.Request.RequestURI != "/admin/login" {
		ctx.Redirect(302, "/admin/login")
	}
}
