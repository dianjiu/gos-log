package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

/**
*CreatedBy:dianjiu
*Time:2017 2021/01/20 14:34
*Project:gos_log
*Discription:使用orm操作postgre,实现对项目日志的增删改查
*o := orm.NewOrm()
*user := UserInfo{Username: "zhangsan", Password: "123456"}
*插入数据
*id, err := o.Insert(&user)
*更新数据
*user.Id = 2
*user.Username = "lisi"
*id, err := o.Update(&user)
*删除数据
*user.Id = 3
*id, err := o.Delete(&user)
*读取
*user.Id = 1
*o.Read(&user)
 */

//AddItem 新增项目日志
func AddItem(item *TItem) (int64, error) {
	o := orm.NewOrm()
	// 插入一条数据，返回自增 id
	id, err := o.Insert(item)
	if err != nil {
		fmt.Println("insert item err : ", err)
	}
	fmt.Println("id :", id)
	return id, err
}

//DeleteItem 根据Id查询项目日志
func DeleteItem(id int64) int64 {
	o := orm.NewOrm()
	item := TItem{}
	item.Id = id
	num, _ := o.Delete(&item)
	return num
}

//UpdateItem 更新项目日志，先查后改
func UpdateItem(item *TItem) (int64, error) {
	o := orm.NewOrm()
	c := TItem{}
	c.Id = item.Id
	err := o.Read(&c)
	if o.Read(&c) == nil {
		c.ClientId = item.ClientId
		c.ItemName = item.ItemName
		c.ItemDesc = item.ItemDesc
		c.LogPath = item.LogPath
		c.LogPrefix = item.LogPrefix
		c.LogSuffix = c.LogSuffix
		c.Status = item.Status
		c.UpdatedTime = time.Now()
		// 修改操作，返回值为受影响的行数
		if num, err := o.Update(&c); err == nil {
			fmt.Println("update return num : ", num)
			return num, err
		}
	}
	return 0, err
}

//ChangeItemStatus 更新项目日志，先查后改
func ChangeItemStatus(id int64) (int64, error) {
	o := orm.NewOrm()
	c := TItem{}
	c.Id = id
	err := o.Read(&c)
	if o.Read(&c) == nil {
		if "1" == c.Status {
			c.Status = "0"
		} else {
			c.Status = "1"
		}
		c.UpdatedTime = time.Now()
		// 修改操作，返回值为受影响的行数
		if num, err := o.Update(&c, "Status", "UpdatedTime"); err == nil {
			fmt.Println("update return num : ", num)
			return num, err
		}
	}
	return 0, err
}

//ReadItem 根据Id查询项目日志
func ReadItem(id int64) (item TItem) {
	o := orm.NewOrm()
	item.Id = id
	err := o.Read(&item)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(item)
	}
	return item
}

// QueryItemsByClientId 根据客户端ID查询所有项目
func QueryItemsByClientId(id int64) (*[]TItem, error) {
	o := orm.NewOrm()
	items := new([]TItem)
	//查找全部
	_, err := o.QueryTable("t_item").Filter("client_id", id).All(items)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

//QueryAllItem 查询所有的项目日志
func QueryAllItem() (*[]TItem, error) {
	o := orm.NewOrm()
	items := new([]TItem)
	//查找全部
	_, err := o.QueryTable("t_item").All(items)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

//QueryPageItem 分页查询所有项目日志
func QueryPageItem(pageNum int, pageSize int) (page Page) {
	o := orm.NewOrm()
	items := new([]TItem)
	o.QueryTable("t_item").Limit(pageSize, (pageNum-1)*pageSize).All(items)
	TotalCount, _ := o.QueryTable("t_item").Count()
	page = PageUtil(int(TotalCount), pageNum, pageSize, items)
	return page
}
