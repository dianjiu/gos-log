package task

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/toolbox"
)

func task() error {
	log.Println("定时检测客户端是否在线 Time:", time.Now().Format("2006-01-02 15:04:05"))
	// TODO 查询所有客户端，协程调用客户端接口
	var server string = "127.0.0.1:2020"
	req := httplib.Get("http://" + server + "/checkOn")
	str, err := req.String()
	if err != nil {
		log.Printf("check logc Online status err: %v", err)
	}
	log.Printf("check logc Online status ok: %v", str)
	// TODO 根据状态，修改数据库标识
	return nil
}

func InitTask() {
	cron := beego.AppConfig.String("cron")
	tk := toolbox.NewTask("task", cron, task)
	toolbox.AddTask("task", tk)
}
