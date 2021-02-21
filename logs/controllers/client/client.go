package client

import (
	"fmt"
	models "logs/models"
	"time"

	"github.com/astaxie/beego"
)

type ClientController struct {
	beego.Controller
}

type Resp struct {
	//必须的大写开头
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data models.Page `json:"data"`
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
	data := Resp{"200", "客户端新增成功", nil}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) Delete() {
	id, _ := this.GetInt64("id")
	models.DeleteClient(id)
	data := Resp{"200", "删除客户端成功", nil}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) Update() {
	client := models.TClient{}
	models.UpdateClient(&client)
	data := Resp{"200", "更新客户端成功", nil}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) QueryAll() {
	clients, _ := models.QueryAllClient()
	this.Data["json"] = &clients
	this.ServeJSON()
}

func (this *ClientController) QueryPage() {
	pageNum, _ := this.GetInt("pageNum")
	pageSize, _ := this.GetInt("pageSize")
	page := models.QueryPageClient(pageNum, pageSize)
	data := Resp{"200", "分页查询客户端成功", page}
	this.Data["json"] = &data
	this.ServeJSON()
}
