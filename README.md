<p align="center">
	<a href="https://goslog.dianjiu.org.cn/"><img src="https://dianjiu.co/goslog.jpg" width="45%"></a>
</p>
<p align="center">
	<strong>Large log retrieval system, so light and fast. </strong>
</p>
<p align="center">
	<a href="https://goslog.dianjiu.org.cn">https://goslog.dianjiu.org.cn/</a>
</p>

-------------------------------------------------------------------------------

[**🌎 Chinese Documentation**](README.zh-CN.md)

-------------------------------------------------------------------------------
## Documentation

[quick starte](https://goslog.dianjiu.org.cn/document/doc01.html)

[version notes](https://goslog.dianjiu.org.cn/version/ver01.html)

## 📦 Use
### Test address（Built on the dev branch ）
[https://goslog.dianjiu.org.cn/test/](https://goslog.dianjiu.org.cn/test/)  
user：admin  
password：admin  

### Production address （Built on the master branch）
[https://goslog.dianjiu.org.cn/prod/](https://goslog.dianjiu.org.cn/prod/)  
user：admin  
password：admin  

## Demo screenshot
### Login page
![login](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105153.jpg)
### Client page
![client](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105216.jpg)
### Item page
![item](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105243.jpg)
### Log search
![logs](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210726105257.jpg)

## Performance Testing

### **System hardware**

| CPU       | Intel® Core™ i5-10210U CPU @ 1.60GHz × 8 |
| --------- | ---------------------------------------- |
| RAM      | 16G                                      |
| ROM      | 512.1 GB                                 |
| Operating system  | Ubuntu 20.04.2 LTS 64位                  |
| GNOME version | 3.36.8                                   |

### **System environment**

| Java   | ORACLE JDK13.0.2     |
| ---------- | -------------------- |
| Go     | GO1.15.7 linux/amd64 |
| Python | Python 3.8.5         |

### Performance comparison

Read single file line by line  17.8G test.log (Single thread reading)

| Language   | test1            | test2  | test3  | test4  | test5  | 总耗时  | 平均耗时 |
| ------ | ---------------- | ------ | ------ | ------ | ------ | ------- | -------- |
| Go     | 32.99s           | 34.24s | 30.33s | 31.21s | 35.70s | 164.16s | 32.83s   |
| Python | Not finished in 32 minutes |        |        |        |        |         |          |
| Java   | 226s             | 206s   | 153s   | 219s   | 183s   | 987s    | 197.4s   |

## Contribute

### Branch description

The source code of GosLog is divided into two branches, with the following functions:

| Branch | Effect                                                  |
|-----------|---------------------------------------------------------------|
| master | The main branch, the branch used by the release version, is consistent with the submission of the central library, and does not receive any pr or modification |
| dev    | Development branch, default is the SNAPSHOT version of the next version, accept modification or pr |

### Provide bug feedback or suggestions

Please indicate the Go version, GosLog version, and related dependency library version you are using when submitting problem feedback.

- [Gitee issue](https://gitee.com/dianjiu/goslog/issues)
- [Github issue](https://github.com/dianjiu/goslog/issues)
- [Gitlab issue](https://gitlab.com/dianjiu/goslog/issues)

### Steps to contribute code

Email to dianjiu@dianjiu.cc to get a developer account

## Star Gos Log
[![Giteye chart](https://chart.giteye.net/gitee/dianjiu/gos-log/96MG4Z3C.png)](https://giteye.net/chart/96MG4Z3C)

[![Stargazers over time](https://starchart.cc/dianjiu/gos-log.svg)](https://starchart.cc/dianjiu/gos-log)

## About author

### Mr. Dianjiu

http://dianjiu.co/








