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
*Discription:使用orm操作postgre,实现对客户端的增删改查
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

//AddClient 新增客户端
func AddClient(client *TClient) (int64, error) {
	o := orm.NewOrm()
	// 插入一条数据，返回自增 id
	id, err := o.Insert(client)
	if err != nil {
		// log.Info("insert client err : ", err)
		fmt.Println("insert client err : ", err)
	}
	fmt.Println("id :", id)
	return id, err
}

//DeleteClient 根据Id查询客户端
func DeleteClient(id int64) int64 {
	o := orm.NewOrm()
	client := TClient{}
	client.Id = id
	num, _ := o.Delete(&client)
	return num
}

//UpdateClient 更新客户端，先查后改
func UpdateClient(client *TClient) (int64, error) {
	o := orm.NewOrm()
	c := TClient{}
	c.Id = client.Id
	err := o.Read(&c)
	if o.Read(&c) == nil {
		c.Ip = client.Ip
		c.Port = client.Port
		c.Vkey = client.Vkey
		c.Info = client.Info
		c.Zip = client.Zip
		c.Online = c.Online
		c.Status = client.Status
		c.UpdatedTime = time.Now()
		// 修改操作，返回值为受影响的行数
		if num, err := o.Update(&c); err == nil {
			fmt.Println("update return num : ", num)
			return num, err
		}
	}
	return 0, err
}

//ChangeClientStatus 更新客户端，先查后改
func ChangeClientStatus(id int64) (int64, error) {
	o := orm.NewOrm()
	c := TClient{}
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

//ReadClient 根据Id查询客户端
func ReadClient(id int64) (client TClient) {
	o := orm.NewOrm()
	client.Id = id
	err := o.Read(&client)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(client)
	}
	return client
}

//QueryAllClient 查询所有的客户端
func QueryAllClient() (*[]TClient, error) {
	o := orm.NewOrm()
	clients := new([]TClient)
	//查找全部
	_, err := o.QueryTable("t_client").All(clients)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return clients, nil
}

//QueryPageClient 分页查询所有客户端
func QueryPageClient(pageNum int, pageSize int) (page Page) {
	o := orm.NewOrm()
	clients := new([]TClient)
	o.QueryTable("t_client").Limit(pageSize, (pageNum-1)*pageSize).All(clients)
	TotalCount, _ := o.QueryTable("t_client").Count()
	page = PageUtil(int(TotalCount), pageNum, pageSize, clients)
	return page
}
