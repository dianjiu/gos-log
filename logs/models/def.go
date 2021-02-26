package models

import "time"

//TClient 客户端
type TClient struct {
	Id          int64     `json:"id" pk:"auto" orm:"column(id)"`
	Ip          string    `json:"ip" orm:"column(ip)"`
	Port        string    `json:"port" orm:"column(port)"`
	Vkey        string    `json:"vkey" orm:"column(vkey)"`
	Info        string    `json:"info" orm:"column(info)"`
	Zip         string    `json:"zip" orm:"column(zip)"`
	Online      string    `json:"online" orm:"column(online)"`
	Status      string    `json:"status" orm:"column(status)"`
	CreatedBy   string    `json:"createdBy" orm:"column(created_by)"`
	CreatedTime time.Time `json:"createdTime" orm:"column(created_time)"`
	UpdatedBy   string    `json:"updatedBy" orm:"column(updated_by)"`
	UpdatedTime time.Time `json:"updatedTime" orm:"column(updated_time)"`
}

//TItem 客户端
type TItem struct {
	Id          int64     `json:"id" pk:"auto" orm:"column(id)"`
	ClientId    int64     `json:"clientId" orm:"column(client_id)"`
	ItemName    string    `json:"itemName" orm:"column(item_name)"`
	ItemDesc    string    `json:"itemDesc" orm:"column(item_desc)"`
	LogPath     string    `json:"logPath" orm:"column(log_path)"`
	LogPrefix   string    `json:"logPrefix" orm:"column(log_prefix)"`
	LogSuffix   string    `json:"logSuffix" orm:"column(log_suffix)"`
	Status      string    `json:"status" orm:"column(status)"`
	CreatedBy   string    `json:"createdBy" orm:"column(created_by)"`
	CreatedTime time.Time `json:"createdTime" orm:"column(created_time)"`
	UpdatedBy   string    `json:"updatedBy" orm:"column(updated_by)"`
	UpdatedTime time.Time `json:"updatedTime" orm:"column(updated_time)"`
}

//Page 分页
type Page struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	FirstPage  bool        `json:"firstPage"`
	LastPage   bool        `json:"lastPage"`
	List       interface{} `json:"list"`
}

//PageUtil 分页工具
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
		List:       list,
	}
}

//DBConfig 数据相关配置
type DBConfig struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	MaxIdleConns int //最大空闲连接
	MaxOpenConns int //最大连接数
}
