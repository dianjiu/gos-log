package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//测试路由
func (this *UserController) Test() {
	this.Ctx.WriteString("这是正则路由 user/test")
}

//Login 用户登陆接口
func (this *UserController) Login() {
	//获取浏览器传入密码
	username := this.Ctx.Input.Param(":username")
	password := this.Ctx.Input.Param(":password")
	//this.Ctx.WriteString("这是正则路由 user/login , username is " + username + " password is " + password)
	//获取配置文件用户信息
	uname := beego.AppConfig.String("username")
	upwd := beego.AppConfig.String("password")
	//判断用户名、密码是否正确
	if uname == username && upwd == password {
		this.Ctx.WriteString("登陆成功")
	} else {
		this.Ctx.WriteString("用户名或密码错误")
	}
}
