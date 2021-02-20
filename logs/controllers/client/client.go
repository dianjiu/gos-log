package client

import (
	"fmt"
	models "logs/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type ClientController struct {
	beego.Controller
}

//Console 控制台
func (this *ClientController) Index() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "client.html"
}

func (this *ClientController) Add() {
	c := models.TClient{}
	c.Ip = "127.0.0.1"
	c.Port = "2020"
	c.Vkey = "123456"
	c.Info = "本地客户端"
	c.Zip = "1"
	c.Online = "0"
	c.Status = "1"
	c.CreatedBy = "admin"
	c.CreatedTime = time.Now()
	c.UpdatedBy = "admin"
	c.UpdatedTime = time.Now()
	id, err := models.AddClient(&c)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	this.Ctx.WriteString("客户端新增成功# " + strconv.Itoa(int(id)))
}
