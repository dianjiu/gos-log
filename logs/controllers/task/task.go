package task

import (
	"encoding/json"
	"errors"
	"log"
	"logs/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/toolbox"
)

type Resp struct {
	//必须的大写开头
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func task() error {
	log.Println("定时检测客户端是否在线 Time:", time.Now().Format("2006-01-02 15:04:05"))
	//TODO 查询所有客户端，协程调用客户端接口
	clients, err := models.QueryAllClient()
	if err != nil {
		log.Printf("orm query all client err: %v", err)
		return errors.New("ORM queryAllClient method is Exception")
	}
	for _, client := range clients {
		logcServer := client.Ip + ":" + client.Port
		req := httplib.Get("http://" + logcServer + "/checkOn")
		str, err := req.String()
		if err != nil {
			log.Printf("check logc Online status err: %v", err)
			//return errors.New("HTTP request detection client status is Exception")
		}
		log.Printf("check logc Online status ok: %v", str)
		var resp Resp
		//字符串转结构体
		json.Unmarshal([]byte(str), &resp)
		if "200" == resp.Code {
			log.Printf("check logc Online status resp code: %v", resp.Code)
			// 根据状态，修改数据库标识
			c := models.TClient{}
			c.Id = client.Id
			c.Online = "1"
			models.ChangeClientOnline(&c)
		} else {
			log.Printf("check logc Online status resp code: %v", resp.Code)
			// 根据状态，修改数据库标识
			c := models.TClient{}
			c.Id = client.Id
			c.Online = "0"
			models.ChangeClientOnline(&c)
		}
	}
	return nil
}

func InitTask() {
	cron := beego.AppConfig.String("cron")
	tk := toolbox.NewTask("task", cron, task)
	toolbox.AddTask("task", tk)
}
