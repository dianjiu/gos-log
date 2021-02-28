package item

import (
	"encoding/json"
	"fmt"
	models "logs/models"
	"time"

	"github.com/astaxie/beego"
)

type ItemController struct {
	beego.Controller
}

type PageResp struct {
	//必须的大写开头
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data models.Page `json:"data"`
}

type ItemResp struct {
	//必须的大写开头
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data models.TItem `json:"data"`
}

//Console 控制台
func (this *ItemController) Index() {
	this.TplName = "item.html"
}

func (this *ItemController) Add() {
	var item models.TItem
	req := this.Ctx.Input.RequestBody
	err := json.Unmarshal(req, &item)
	if err == nil {
		item.CreatedBy = "admin"
		item.CreatedTime = time.Now()
		item.UpdatedBy = "admin"
		item.UpdatedTime = time.Now()
		id, err := models.AddItem(&item)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)
		data := ItemResp{"200", "项目日志新增成功", models.TItem{}}
		this.Data["json"] = &data
		this.ServeJSON()
	}
}

func (this *ItemController) Delete() {
	id, _ := this.GetInt64("id")
	models.DeleteItem(id)
	data := ItemResp{"200", "删除项目日志成功", models.TItem{}}
	this.Data["json"] = &data
	this.ServeJSON()
}
func (this *ItemController) Query() {
	id, _ := this.GetInt64("id")
	// id, _ := this.GetInt("id")
	item := models.ReadItem(id)
	data := ItemResp{"200", "查询项目日志成功", item}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ItemController) Update() {
	// client := models.TClient{}
	var item models.TItem
	req := this.Ctx.Input.RequestBody
	err := json.Unmarshal(req, &item)
	if err == nil {
		models.UpdateItem(&item)
		data := ItemResp{"200", "更新项目日志成功", models.TItem{}}
		this.Data["json"] = &data
		this.ServeJSON()
	}
}

func (this *ItemController) ChangeItemStatus() {
	id, _ := this.GetInt64("id")
	models.ChangeItemStatus(id)
	data := ItemResp{"200", "更新项目日志成功", models.TItem{}}
	this.Data["json"] = &data
	this.ServeJSON()
}

func (this *ItemController) QueryAll() {
	items, _ := models.QueryAllItem()
	this.Data["json"] = &items
	this.ServeJSON()
}

func (this *ItemController) QueryPage() {
	pageNum, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("limit")
	page := models.QueryPageItem(pageNum, pageSize)
	data := PageResp{"200", "分页查询项目日志成功", page}
	this.Data["json"] = &data
	this.ServeJSON()
}
