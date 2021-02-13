package routers

import (
	"github.com/astaxie/beego"
	"logc/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
