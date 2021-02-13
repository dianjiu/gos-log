package routers

import (
	"github.com/astaxie/beego"
	"logs/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
