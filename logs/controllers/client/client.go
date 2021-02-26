package client

import (
	"encoding/json"
	"fmt"
	models "logs/models"
	"time"

	"github.com/astaxie/beego"
)

type ClientController struct {
	beego.Controller
}

type PageResp struct {
	//必须的大写开头
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data models.Page `json:"data"`
}

type ClientResp struct {
	//必须的大写开头
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data models.TClient `json:"data"`
}

//Console 控制台
func (this *ClientController) Index() {
	this.TplName = "client.html"
}

func (this *ClientController) Add() {
	var client models.TClient
	req := this.Ctx.Input.RequestBody
	err := json.Unmarshal(req, &client)
	if err == nil {
		client.Online = "0"
		client.CreatedBy = "admin"
		client.CreatedTime = time.Now()
		client.UpdatedBy = "admin"
		client.UpdatedTime = time.Now()
		id, err := models.AddClient(&client)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)
		data := ClientResp{"200", "客户端新增成功", models.TClient{}}
		this.Data["json"] = &data
		this.ServeJSON()
	}
}

func (this *ClientController) Delete() {
	id, _ := this.GetInt64("id")
	models.DeleteClient(id)
	data := ClientResp{"200", "删除客户端成功", models.TClient{}}
	this.Data["json"] = &data
	this.ServeJSON()
}
func (this *ClientController) Query() {
	id, _ := this.GetInt64("id")
	// id, _ := this.GetInt("id")
	client := models.ReadClient(id)
	data := ClientResp{"200", "查询客户端成功", client}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ClientController) Update() {
	// client := models.TClient{}
	var client models.TClient
	req := this.Ctx.Input.RequestBody
	err := json.Unmarshal(req, &client)
	if err == nil {
		models.UpdateClient(&client)
		data := ClientResp{"200", "更新客户端成功", models.TClient{}}
		this.Data["json"] = &data
		this.ServeJSON()
	}
}

func (this *ClientController) ChangeStatus() {
	id, _ := this.GetInt64("id")
	models.ChangeStatus(id)
	data := ClientResp{"200", "更新客户端成功", models.TClient{}}
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
	data := PageResp{"200", "分页查询客户端成功", page}
	this.Data["json"] = &data
	this.ServeJSON()
}
