package client

import "github.com/astaxie/beego"

type ClientController struct {
	beego.Controller
}

//Console 控制台
func (this *ClientController) Index() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "client.html"
}
