package register

import (
	"log"
	"os"

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

func (this *RegisterController) Register() {
	server := beego.AppConfig.String("logs")
	vKey := beego.AppConfig.String("key")
	registerLocalIp(server, vKey)
	// log.Printf("Local client registered successfully.")
}

func registerLocalIp(server string, vKey string) {
	//通过Http调用客户端
	req := httplib.Post("http://" + server + "/client/register").Debug(true)
	req.JSONBody(map[string]interface{}{"key": vKey})
	_, err := req.String()
	if err != nil {
		log.Printf("Local client registered error.")
	} else {
		log.Printf("Local client registered successfully.")
	}
	os.Exit(0)
}
