<p align="center">
	<a href="https://goslog.dianjiu.org.cn/"><img src="https://dianjiu.co/goslog.jpg" width="45%"></a>
</p>
<p align="center">
	<strong>轻量级高性能的大日志检索系统</strong>
</p>
<p align="center">
	<a href="https://goslog.dianjiu.org.cn">https://goslog.dianjiu.org.cn/</a>
</p>

-------------------------------------------------------------------------------

[**🌎 English Documentation**](README.md)

-------------------------------------------------------------------------------
## 文档

[快速开始](https://goslog.dianjiu.org.cn/document/doc01.html)

[版本日志](https://goslog.dianjiu.org.cn/version/ver01.html)

## 体验
### 测试地址（基于dev分支构建）
[https://goslog.dianjiu.org.cn/test/](https://goslog.dianjiu.org.cn/test/)  
账号：admin  
密码：admin

### 生产地址（基于master分支构建）
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

## 贡献

### 分支说明

Gos Log的源码分为两个分支，功能如下：

| 分支       | 作用                                                          |
|-----------|---------------------------------------------------------------|
| master | 主分支，release版本使用的分支，与中央库提交一致，不接收任何pr或修改 |
| dev    | 开发分支，默认为下个版本的SNAPSHOT版本，接受修改或pr                 |

### 提供bug反馈或建议

提交问题反馈请说明正在使用的Go版本、GosLog版本和相关依赖库版本。

- [Gitee issue](https://gitee.com/dianjiu/goslog/issues)
- [Github issue](https://github.com/dianjiu/goslog/issues)
- [Gitlab issue](https://gitlab.com/dianjiu/goslog/issues)

### 贡献代码的步骤

邮件联系 dianjiu@dianjiu.cc 获取开发者账号

## Star Gos Log
[![Giteye chart](https://chart.giteye.net/gitee/dianjiu/gos-log/96MG4Z3C.png)](https://giteye.net/chart/96MG4Z3C)

[![Stargazers over time](https://starchart.cc/dianjiu/gos-log.svg)](https://starchart.cc/dianjiu/gos-log)


## 关于作者

### 点九先生

http://dianjiu.co/
