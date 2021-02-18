package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//Test 测试路由
func (this *UserController) Test() {
	this.Ctx.WriteString("这是正则路由 user/test")
}

//Console 控制台
func (this *UserController) Console() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "console.html"
}

//Login 用户登陆接口
func (this *UserController) Login() {
	if this.Ctx.Input.IsGet() {
		// 获取 session
		userName := this.GetSession("username")
		userPwd := this.GetSession("password")
		_, nameOk := userName.(string)
		_, pwdOk := userPwd.(string)
		if nameOk && pwdOk {
			// 重定向
			this.Redirect("/admin/index", 302)
		} else {
			this.TplName = "login.html"
		}
	} else {
		// 获取请求参数
		userName := this.GetString("username")
		userPwd := this.GetString("password")
		//获取配置文件用户信息
		uname := beego.AppConfig.String("username")
		upwd := beego.AppConfig.String("password")
		// 表单验证
		//判断用户名、密码是否正确
		if uname == userName && upwd == userPwd {
			//this.Ctx.WriteString("登陆成功")
			// 设置 session
			this.SetSession("username", userName)
			this.SetSession("password", userPwd)
			this.Redirect("/admin/index", 302)
		} else {
			//this.Ctx.WriteString("用户名或密码错误")
			// 重定向
			this.Redirect("/admin/login", 302)
		}
	}

}

// Index 首页
func (this *UserController) Index() {
	user_name := this.GetSession("username")
	this.Data["username"] = user_name
	this.TplName = "index.html"
}

// Exit 推出登陆
func (this *UserController) Exit() {
	// 清空 session ，清空后 key 对应的 session value 是 nil
	this.DelSession("username")
	this.DelSession("password")
	this.Data["json"] = nil
	this.ServeJSON()
	this.Redirect("/admin/login", 302)
}
