## gos-log
基于Go语言的轻量级高性能的大日志检索系统

## 开源地址

### gos-log

[https://gitee.com/dianjiu/gos-log](https://gitee.com/dianjiu/gos-log)  

[https://github.com/dianjiu/gos-log](https://github.com/dianjiu/gos-log)  

### gos-log-vue

[https://gitee.com/dianjiu/gos-log-vue](https://gitee.com/dianjiu/gos-log-vue)  

[https://github.com/dianjiu/gos-log-vue](https://github.com/dianjiu/gos-log-vue)  

## 官方地址
[http://goslog.dianjiu.org.cn](http://goslog.dianjiu.org.cn)  

### 测试地址（最新功能体验）
[https://goslog.dianjiu.org.cn/test/](https://goslog.dianjiu.org.cn/test/)  
账号：admin  
密码：admin  

### 生产地址（稳定版本体验）
[https://goslog.dianjiu.org.cn/prod/](https://goslog.dianjiu.org.cn/prod/)  
账号：admin  
密码：admin  

## 演示截图
### 登陆页
![login](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105153.jpg)
### 客户端
![client](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105216.jpg)
### 项目管理
![item](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105243.jpg)
### 日志查找
![logs](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105257.jpg)

## 性能测试

### **系统硬件**

| CPU       | Intel® Core™ i5-10210U CPU @ 1.60GHz × 8 |
| --------- | ---------------------------------------- |
| 内存      | 16G                                      |
| 硬盘      | 512.1 GB                                 |
| 操作系统  | Ubuntu 20.04.2 LTS 64位                  |
| GNOME版本 | 3.36.8                                   |

### **系统环境**

| Java环境   | ORACLE JDK13.0.2     |
| ---------- | -------------------- |
| Go环境     | GO1.15.7 linux/amd64 |
| Python环境 | Python 3.8.5         |

### 性能对比

单文件逐行读取  17.8G test.log (单线程读取)

| 语言   | test1            | test2  | test3  | test4  | test5  | 总耗时  | 平均耗时 |
| ------ | ---------------- | ------ | ------ | ------ | ------ | ------- | -------- |
| Go     | 32.99s           | 34.24s | 30.33s | 31.21s | 35.70s | 164.16s | 32.83s   |
| Python | 32分钟还没执行完 |        |        |        |        |         |          |
| Java   | 226s             | 206s   | 153s   | 219s   | 183s   | 987s    | 197.4s   |

## 版本日志
[https://goslog.dianjiu.org.cn/version/ver01.html](https://goslog.dianjiu.org.cn/version/ver01.html)

## 快速上手
[https://goslog.dianjiu.org.cn/document/doc01.html](https://goslog.dianjiu.org.cn/document/doc01.html)


## 关于点九

### 点九先生

http://dianjiu.co/

### 个人邮箱

dianjiu@dianjiu.cc

## 项目致谢





