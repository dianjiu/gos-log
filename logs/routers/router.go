package routers

import (
	"logs/controllers/admin"
	"logs/controllers/client"
	"logs/controllers/item"
	"logs/controllers/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token", "X-Requested-With", "withCredentials"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	// 验证用户是否已经登录
	// beego.InsertFilter("/*", beego.BeforeExec, FilterUser)
	beego.Router("/", &admin.UserController{}, "*:Login")
	beego.Router("/admin/test", &admin.UserController{}, "*:Test")
	beego.Router("/admin/console", &admin.UserController{}, "*:Console")
	//beego.Router("/admin/index", &admin.UserController{}, "*:Index")
	beego.Router("/admin/login", &admin.UserController{}, "*:Login")
	beego.Router("/admin/exit", &admin.UserController{}, "*:Exit")
	//beego.Router("/client/index", &client.ClientController{}, "*:Index")
	beego.Router("/client/register", &client.ClientController{}, "*:Register")
	beego.Router("/client/add", &client.ClientController{}, "*:Add")
	beego.Router("/client/delete", &client.ClientController{}, "*:Delete")
	beego.Router("/client/update", &client.ClientController{}, "*:Update")
	beego.Router("/client/changeStatus", &client.ClientController{}, "*:ChangeClientStatus")
	beego.Router("/client/query", &client.ClientController{}, "*:Query")
	beego.Router("/client/queryAll", &client.ClientController{}, "*:QueryAll")
	beego.Router("/client/queryPage", &client.ClientController{}, "*:QueryPage")
	//beego.Router("/client/index", &client.ClientController{}, "*:Index")
	beego.Router("/item/add", &item.ItemController{}, "*:Add")
	beego.Router("/item/delete", &item.ItemController{}, "*:Delete")
	beego.Router("/item/update", &item.ItemController{}, "*:Update")
	beego.Router("/item/changeStatus", &item.ItemController{}, "*:ChangeItemStatus")
	beego.Router("/item/query", &item.ItemController{}, "*:Query")
	beego.Router("/item/queryAll", &item.ItemController{}, "*:QueryAll")
	beego.Router("/item/queryPage", &item.ItemController{}, "*:QueryPage")
	//beego.Router("/item/index", &item.ItemController{}, "*:Index")
	//beego.Router("/logs/index", &logs.LogsController{}, "*:Index")
	beego.Router("/logs/query", &logs.LogsController{}, "*:Query")
	beego.Router("/logs/queryClients", &logs.LogsController{}, "*:QueryClient")
	beego.Router("/logs/queryItems", &logs.LogsController{}, "*:QueryItem")
}

/* var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("username").(string)

	if !ok && ctx.Request.RequestURI != "/admin/login" {
		ctx.Redirect(302, "/admin/login")
	}
} */
