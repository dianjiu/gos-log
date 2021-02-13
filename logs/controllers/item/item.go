package item

import "github.com/astaxie/beego"

type ItemController struct {
	beego.Controller
}

//Console 控制台
func (this *ItemController) Index() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "item.html"
}
