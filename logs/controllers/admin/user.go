package admin

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

type UserResp struct {
	//必须的大写开头
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data User   `json:"data"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Test 测试路由
func (this *UserController) Test() {
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")
	// TODO 代办测试
	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
	this.Ctx.WriteString("这是正则路由 user/test")
}

//Console 控制台
func (this *UserController) Console() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "console.html"
}

//Login 用户登陆接口
func (this *UserController) Login() {
	//先获取 session,判断用户是不是已经登录了
	userName := this.GetSession("username")
	userPwd := this.GetSession("password")
	unameStr, nameOk := userName.(string)
	upwdStr, pwdOk := userPwd.(string)
	logs.Info("username= %s password= %s", unameStr, upwdStr)
	if nameOk && pwdOk {
		// 重定向
		// this.Redirect("/admin/index", 302)
		data := UserResp{"200", "用户已经登陆", User{"admin", ""}}
		this.Data["json"] = &data
		this.ServeJSON()
	}
	//用户没有登录，获取请求参数
	/* userName := this.GetString("username")
	userPwd := this.GetString("password") */
	var user User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err == nil {
		//获取配置文件用户信息
		uname := beego.AppConfig.String("username")
		upwd := beego.AppConfig.String("password")
		// 表单验证
		//判断用户名、密码是否正确
		if uname == user.Username && upwd == user.Password {
			//this.Ctx.WriteString("登陆成功")
			// 设置 session
			this.SetSession("username", userName)
			this.SetSession("password", userPwd)
			// this.Redirect("/admin/index", 302)
			data := UserResp{"200", "登陆成功", User{uname, ""}}
			this.Data["json"] = &data
			this.ServeJSON()
		} else {
			//this.Ctx.WriteString("用户名或密码错误")
			// 重定向
			// this.Redirect("/admin/login", 302)
			data := UserResp{"302", "用户名或密码错误", User{}}
			this.Data["json"] = &data
			this.ServeJSON()
		}
	}
}

// Index 首页
/*func (this *UserController) Index() {
	user_name := this.GetSession("username")
	this.Data["username"] = user_name
	this.TplName = "index.html"
}*/

// Exit 推出登陆
func (this *UserController) Exit() {
	// 清空 session ，清空后 key 对应的 session value 是 nil
	this.DelSession("username")
	this.DelSession("password")
	/* this.Data["json"] = nil
	this.ServeJSON()
	this.Redirect("/admin/login", 302) */
	data := UserResp{"200", "退出成功", User{}}
	this.Data["json"] = &data
	this.ServeJSON()
}
