package task

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func task() error {
	log.Println("定时检测客户端是否在线 Time:", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func InitTask() {
	cron := beego.AppConfig.String("cron")
	tk := toolbox.NewTask("task", cron, task)
	toolbox.AddTask("task", tk)
}
