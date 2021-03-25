package file

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

type FileReq struct {
	Path string `json:"path"`
	Key  string `json:"key"`
	Line int64  `json:"line"`
}

func (this *FileController) Query() {
	var file FileReq
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &file)
	// line, err := strconv.ParseInt(file.Line, 10, 64)
	if err == nil {
		file := ReadString(file.Path, file.Key, file.Line)
		// file := ReadString("c:/logs/server.2021-01-04.log", "b77ee6ba0b4cdd28", 1000)
		this.Ctx.Output.Download(file)
	}
}

//ReadString
func ReadString(filename string, key string, line int64) (file string) {
	startTime := time.Now()
	defer func(startTime time.Time) {
		fmt.Printf("共耗时：%s\n", time.Now().Sub(startTime).String())
	}(startTime)
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	var lineBegin int64 = 0
	var lineFirst int64 = 0
	var lineOver int64 = 0
	// 获取临时目录
	temppath := beego.AppConfig.String("temppath")
	// 把查找到的日志放到临时目录下 文件名以关键字命名
	dstFile, err := os.OpenFile(temppath+key+".log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	//首次便利记录出现行号
	for {
		str, err := r.ReadString('\n')
		lineBegin = lineBegin + 1
		if err != nil {
			break
		}
		if strings.Contains(str, key) {
			// fmt.Println(str)
			if lineFirst == 0 && lineOver == 0 {
				lineFirst = lineBegin
				lineOver = lineFirst + line
				fmt.Printf("lineFirst:%d\n", lineFirst)
				fmt.Printf("lineBegin:%d\n", lineBegin)
				fmt.Printf("lineOver:%d\n", lineOver)
			}
		}
		// 暂不考虑lineOver 小于 over 的时候
		if lineBegin <= lineOver {
			bufWriter.WriteString(str)
		}
		if lineOver > 0 && lineBegin > lineOver {
			break
		}
	}
	bufWriter.Flush()
	dstFile.Close()
	fmt.Printf("查找耗时：%s\n", time.Now().Sub(startTime).String())
	//二次遍历输出文件
	/* startTimeTwo := time.Now()
	var start int64 = 0
	// 获取临时目录
	temppath := beego.AppConfig.String("temppath")
	// 把查找到的日志放到临时目录下 文件名以关键字命名
	dstFile, err := os.OpenFile(temppath+key+".log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	f1, _ := os.Open(filename)
	defer f1.Close()
	r1 := bufio.NewReader(f1)
	for {
		str, err := r1.ReadString('\n')
		start = start + 1
		if err != nil {
			break
		}
		// 暂不考虑lineOver 小于 over 的时候
		if start >= lineBegin && start <= lineOver {
			bufWriter.WriteString(str)
		}
		if start > lineOver {
			break
		}
	}
	bufWriter.Flush()
	dstFile.Close()
	fmt.Printf("写出耗时：%s\n", time.Now().Sub(startTimeTwo).String()) */
	// 压缩
	startTimeThree := time.Now()
	ip := GetLocalIPv4()
	var src string = temppath + key + ".log"
	var dst string = temppath + ip + ".zip"
	Zip(dst, src)
	fmt.Printf("压缩耗时：%s\n", time.Now().Sub(startTimeThree).String())
	return dst
}

//GetLocalIPv4 获取本机的IPv4地址
func GetLocalIPv4() (ip string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println("本机IP：" + ipnet.IP.String())
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}

// Zip 压缩文件
func Zip(dst, src string) (err error) {
	s, _ := ioutil.ReadFile(src)
	var res bytes.Buffer
	gz := gzip.NewWriter(&res)
	gz.Write(s)
	gz.Flush()
	ioutil.WriteFile(dst, res.Bytes(), 0777)
	return nil
}

// UnZip 解压缩
func UnZip(dst, src string) (err error) {
	s, _ := ioutil.ReadFile(dst)
	var res bytes.Buffer
	binary.Write(&res, binary.LittleEndian, s)
	gz, _ := gzip.NewReader(&res)
	data, _ := ioutil.ReadAll(gz)
	defer gz.Close()
	ioutil.WriteFile(src, data, 0777)
	return nil
}
