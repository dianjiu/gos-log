package register

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type RegisterController struct {
	beego.Controller
}

type Resp struct {
	//必须的大写开头
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Register Local client registered by local app.conf
func (this *RegisterController) Register() {
	server := beego.AppConfig.String("logs")
	vKey := beego.AppConfig.String("key")
	RegisterLocalIp(server, vKey)
	// log.Printf("Local client registered successfully.")
}

// CheckOnline Online interface for logs detection
func (this *RegisterController) CheckOnline() {
	data := Resp{"200", "客户端在线"}
	this.Data["json"] = &data
	this.ServeJSON()
}

// RegisterLocalIp Automatically register the client at startup
func RegisterLocalIp(server string, vKey string) {
	//通过Http调用客户端
	req := httplib.Post("http://" + server + "/client/register").Debug(true)
	req.JSONBody(map[string]interface{}{"key": vKey})
	_, err := req.String()
	log.Printf("logc register url=%v param=%v errMsg=%v\n", "http://"+server+"/client/register", vKey, err)
	if err != nil {
		log.Printf("Local client registered error.")
	} else {
		log.Printf("Local client registered successfully.")
	}
}
