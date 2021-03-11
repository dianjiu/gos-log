package main

import (
	models "logs/models"
	_ "logs/routers"

	"github.com/astaxie/beego"
)

func init() {
	//初始化数据链接
	db := models.DBConfig{
		Host:         "127.0.0.1",
		Port:         "3306",
		Username:     "root",
		Password:     "123456",
		Database:     "demo",
		MaxIdleConns: 10,
		MaxOpenConns: 50,
	}
	models.NewDef(&db)
}

func main() {
	beego.Run()
}
