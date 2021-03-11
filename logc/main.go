package main

import (
	_ "logc/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 定义变量，用于接收命令行的参数值
	// var server string
	// var vkey string
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	// flag.StringVar(&server, "s", "127.0.0.1:2021", "ip+port")
	// flag.StringVar(&vkey, "v", "123456", "密钥")
	// 解析命令行参数写入注册的flag里
	// flag.Parse()
	// 输出结果
	// fmt.Printf("logc register -server=%v -vkey=%v\n", server, vkey)
	// register.Register(server, vkey)
	beego.Run()
}
