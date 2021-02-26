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
	this.TplName = "client.html"
}

func (this *ClientController) Add() {
	c := models.TClient{}
	c.Ip = this.GetString("ip")
	c.Port = this.GetString("port")
	c.Vkey = this.GetString("vkey")
	c.Info = this.GetString("info")
	c.Zip = this.GetString("zip")
	c.Online = this.GetString("online")
	c.Status = this.GetString("status")
	c.CreatedBy = "admin"
	c.CreatedTime = time.Now()
	c.UpdatedBy = "admin"
	c.UpdatedTime = time.Now()
	id, err := models.AddClient(&c)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	data := Resp{"200", "客户端新增成功", models.Page{}}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) Delete() {
	id, _ := this.GetInt64("id")
	models.DeleteClient(id)
	data := Resp{"200", "删除客户端成功", models.Page{}}
	this.Data["json"] = &data
	this.ServeJSON()
}
func (this *ClientController) Query() {
	id, _ := this.GetInt64("id")
	models.ReadClient(id)
	data := Resp{"200", "查询客户端成功", models.Page{}}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) Update() {
	client := models.TClient{}
	models.UpdateClient(&client)
	data := Resp{"200", "更新客户端成功", models.Page{}}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) QueryAll() {
	clients, _ := models.QueryAllClient()
	this.Data["json"] = &clients
	this.ServeJSON()
}

func (this *ClientController) QueryPage() {
	pageNum, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("limit")
	page := models.QueryPageClient(pageNum, pageSize)
	data := Resp{"200", "分页查询客户端成功", page}
	this.Data["json"] = &data
	this.ServeJSON()
}
