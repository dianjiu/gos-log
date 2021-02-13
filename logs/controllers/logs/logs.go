package logs

import "github.com/astaxie/beego"

type LogsController struct {
	beego.Controller
}

//Console 控制台
func (this *LogsController) Index() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "logs.html"
}
