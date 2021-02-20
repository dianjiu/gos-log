package models

import (
	"fmt"

	log "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

/**
*CreatedBy:dianjiu
*Time:2017 2021/01/20 14:34
*Project:gos_log
*Discription:使用orm操作postgre,实现对客户端的增删改查
 */
type ClientManager struct {
	DBConf *DBConfig
}

/*
 * clientmanger构造器
 */
func NewClientManager(dbConf *DBConfig) *ClientManager {
	mgr := &ClientManager{
		DBConf: dbConf,
	}
	//初始化orm
	mgr.initDB()
	return mgr
}

/**
  初始化db，注册默认数据库，同时将实体模型也注册上去
*/
func (mgr *ClientManager) initDB() {
	// 是否开启调试模式 调试模式下会打印出sql语句
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DRPostgres)
	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", mgr.DBConf.Host, mgr.DBConf.Port, mgr.DBConf.Username, mgr.DBConf.Password, mgr.DBConf.Database)
	log.Info("datasource=[%s]", ds)
	err := orm.RegisterDataBase("default", "postgres", ds, mgr.DBConf.MaxIdleConns, mgr.DBConf.MaxOpenConns)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(TClient))
}

//AddClient 新增客户端
func AddClient(client *TClient) (int64, error) {
	o := orm.NewOrm()
	// 插入一条数据，返回自增 id
	id, err := o.Insert(client)
	if err != nil {
		fmt.Println("insert client err : ", err)
	}
	fmt.Println("id :", id)
	return id, err
}

/**
 *通过orm对象来进行数据库的操作，这种情况是必须要知道主键
 */
func (mgr *ClientManager) OpreateClientWithOrmObjct(id int64) {
	orm.Debug = true
	o := orm.NewOrm()
	var clients = new(TClient) //注意使用new来创建变量，否则就会报错
	clients.Id = id
	fmt.Print("******开始读取数据库*******")
	o.Read(clients)
	fmt.Print(clients)
	fmt.Print("******开始更新数据库*******")
	clients.Ip = "127.0.0.1"
	o.Update(clients)
	o.Read(clients)
	fmt.Printf("更新后的数据库为%v", clients)
	o.Delete(clients)
}

/**
根据某些字段来read 1：采用queryTable方法来查询 2：采用Raw执行sql语句
*/
func (mgr *ClientManager) GetClientsByIdWithQueryTable(id string) (*[]TClient, error) {
	orm.Debug = true
	o := orm.NewOrm()
	client := new([]TClient)
	//查找全部
	_, err := o.QueryTable("client").Filter("Id", id).All(client)
	//查找单个
	//err := o.QueryTable("client").Filter("id",key).One(client)
	//使用sql语句进行查询
	//err:=o.Raw("select * from t_client where Id = ?",id).QueryRow(client)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil
}

func (mgr *ClientManager) getClinetsWithPage(ps int, pageSize int) (page Page) {
	o := orm.NewOrm()
	client := new([]TClient)
	o.QueryTable("client").Limit(pageSize, (ps-1)*pageSize).All(client)
	TotalCount, _ := o.QueryTable("client").Count()
	page.TotalCount = int(TotalCount)
	page.PageSize = pageSize
	page.List = client
	fmt.Println(client)
	return page
}
