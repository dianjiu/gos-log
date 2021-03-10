package register

import (
	"log"
	"os"

	"github.com/astaxie/beego/httplib"
)

type Resp struct {
	//必须的大写开头
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func Register(server string, vKey string) {
	log.Printf("Local client registered successfully.")
}

func RegisterLocalIp(server string, vKey string) {
	var result Resp
	//通过Http调用客户端
	req := httplib.Post(server + "/client/register").Debug(true)
	req.JSONBody(map[string]interface{}{"key": vKey})
	req.ToJson(&result)
	if result.Code == "200" {
		log.Printf("Local client registered successfully.")
	}
	os.Exit(0)
}
