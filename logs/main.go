package main

import (
	"logs/controllers/task"
	models "logs/models"
	_ "logs/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	sqlhost := beego.AppConfig.String("sqlhost")
	sqlport := beego.AppConfig.String("sqlport")
	sqluser := beego.AppConfig.String("sqluser")
	sqlpwd := beego.AppConfig.String("sqlpwd")
	database := beego.AppConfig.String("database")
	maxIdleConns, _ := beego.AppConfig.Int("maxIdleConns")
	maxOpenConns, _ := beego.AppConfig.Int("maxOpenConns")
	//初始化数据链接
	db := models.DBConfig{
		Host:         sqlhost,
		Port:         sqlport,
		Username:     sqluser,
		Password:     sqlpwd,
		Database:     database,
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
		/* Host:         "127.0.0.1",
		Port:         "3306",
		Username:     "root",
		Password:     "123456",
		Database:     "demo",
		MaxIdleConns: 10,
		MaxOpenConns: 50, */
	}
	models.NewDef(&db)
	//初始化定时任务
	task.InitTask()
}

func main() {
	//定时任务启动
	toolbox.StartTask()
	defer toolbox.StopTask()
	beego.Run()
}
