package logs

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	models "logs/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type LogsController struct {
	beego.Controller
}

type LogsReq struct {
	Client int64  `json:"client"`
	Item   int64  `json:"item"`
	Date   string `json:"date"`
	Key    string `json:"key"`
	Line   int64  `json:"line"`
}

type LogsResp struct {
	//必须的大写开头
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var wg sync.WaitGroup

//Console 控制台
func (this *LogsController) Index() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	//this.TplName = "logs.html"
}

func (this *LogsController) Query() {
	var logs LogsReq
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &logs)
	// 获取临时目录
	temppath := beego.AppConfig.String("temppath")
	isExists, _ := pathExists(temppath + "/" + logs.Key)
	if isExists == false {
		createDir(temppath + "/" + logs.Key)
	}
	// line, err := strconv.ParseInt(file.Line, 10, 64)
	if err == nil {
		// client为0的话，需要查询所有的在线客户端
		if logs.Client != 0 {
			// 根据客户端ID查在线的客户端服务IP:Port
			client := models.ReadClient(logs.Client)
			url := "http://" + client.Ip + ":" + client.Port + "/file/query"
			// url := "http://localhost:2020/file/query"
			// 根据项目ID查日志路径、前缀、后缀，拼接日志全路径
			item := models.ReadItem(logs.Item)
			path := item.LogPath + item.LogPrefix + logs.Date + item.LogSuffix + ".log"
			// path := "C:\\logs\\server.2021-01-04.log"
			//通过Http调用客户端
			req := httplib.Post(url).Debug(true)
			req.JSONBody(map[string]interface{}{"path": path, "key": logs.Key, "line": logs.Line})
			req.ToFile(temppath + logs.Key + "/" + client.Ip + ".zip")
		} else {
			//获取所有客户端遍历
			clients, _ := models.QueryAllClient()
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "多客户端并发查询开始了")
			done := make(chan bool)
			data := make(chan int)
			wg.Add(len(clients))
			go func() {
				wg.Wait()
				close(data)
			}()
			go outClient(data, done)
			// for _, client := range clients {
			for i := 0; i < len(clients); i++ {
				//wg.Add(1)
				go gotoClient(data, done, i, clients[i], logs, temppath)
				// 根据客户端ID查在线的客户端服务IP:Port
				/* url := "http://" + client.Ip + ":" + client.Port + "/file/query"
				// url := "http://localhost:2020/file/query"
				// 根据项目ID查日志路径、前缀、后缀，拼接日志全路径
				item := models.ReadItem(logs.Item)
				path := item.LogPath + item.LogPrefix + logs.Date + item.LogSuffix + ".log"
				// path := "C:\\logs\\server.2021-01-04.log"
				//通过Http调用客户端
				// go getLog(url, path, logs.Key, temppath, client.Ip, logs.Line)
				req := httplib.Post(url).Debug(true)
				req.JSONBody(map[string]interface{}{"path": path, "key": logs.Key, "line": logs.Line})
				req.ToFile(temppath + "/" + logs.Key + "/" + client.Ip + ".zip") */
			}
			<-done
		}
		Zip(temppath+"/"+logs.Key+".zip", temppath+"/"+logs.Key+"/")
		defer func() {
			//返回后清理压缩文件
			os.Remove(temppath + "/" + logs.Key + ".zip")
			os.RemoveAll(temppath + "/" + logs.Key + "/")
		}()
		// this.Ctx.WriteString("好了")
		defer func() {
			//返回后清理压缩文件
			os.Remove(temppath + "/" + logs.Key + ".zip")
		}()
		this.Ctx.Output.Download(temppath + "/" + logs.Key + ".zip")
	}
}

func outClient(data chan int, done chan bool) {
	for x := range data {
		i := strconv.Itoa(x)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "调用客户端"+i+"结束")
	}
	done <- true
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "多客户端并发查询结束了")
}

func gotoClient(data chan int, done chan bool, i int, client models.TClient, logs LogsReq, temppath string) {
	defer wg.Done()
	x := strconv.Itoa(i)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "调用客户端"+x+"开始")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), client)
	// 根据客户端ID查在线的客户端服务IP:Port
	url := "http://" + client.Ip + ":" + client.Port + "/file/query"
	// 根据项目ID查日志路径、前缀、后缀，拼接日志全路径
	item := models.ReadItem(logs.Item)
	path := item.LogPath + item.LogPrefix + logs.Date + item.LogSuffix + ".log"
	//通过Http调用客户端
	req := httplib.Post(url).Debug(true)
	req.JSONBody(map[string]interface{}{"path": path, "key": logs.Key, "line": logs.Line})
	req.ToFile(temppath + logs.Key + "/" + client.Ip + ".zip")
	data <- i
}

//getLog
func getLog(url, path, key, temppath, ip string, line int64) {
	//通过Http调用客户端
	req := httplib.Post(url).Debug(true)
	req.JSONBody(map[string]interface{}{"path": path, "key": key, "line": line})
	req.ToFile(temppath + "/" + key + "/" + ip + ".zip")
}

func (this *LogsController) QueryClient() {
	clients, _ := models.QueryAllClient()
	data := LogsResp{"200", "查询客户端列表成功", clients}
	this.Data["json"] = &data
	this.ServeJSON()

}

func (this *LogsController) QueryItem() {
	clientId, _ := this.GetInt64("client_id")
	items, _ := models.QueryItemsByClientId(clientId)
	data := LogsResp{"200", "根据客户端ID查询项目日志成功", items}
	this.Data["json"] = &data
	this.ServeJSON()
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { // 文件或者目录存在
		// log.Fatal(err)
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// createDir 创建文件夹
func createDir(folderPath string) string {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// removeDir 删除指定目录下所有文件
func removeDir(folderPath string) {
	os.Remove(folderPath)
}

// removeAllDir 删除指定目录
func removeAllDir(folderPath string) {
	os.RemoveAll(folderPath)
}

// Zip 压缩文件
func Zip(dst, src string) (err error) {
	// 创建准备写入的文件
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
		//压缩成功后 删除原文件
		os.Remove(src)
	}()

	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		// 通过文件信息，创建 zip 的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		// 替换文件信息中的文件名
		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		// 这步开始没有加，会发现解压的时候说它不是个目录
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 Write 结构
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		// 将打开的文件 Copy 到 w
		n, err := io.Copy(w, fr)
		if err != nil {
			return
		}
		// 输出压缩的内容
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)

		return nil
	})
}

// UnZip 解压缩
func UnZip(dst, src string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return
	}

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		// 创建要写出的文件对应的 Write
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		n, err := io.Copy(fw, fr)
		if err != nil {
			return err
		}

		// 将解压的结果输出
		fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

		// 因为是在循环中，无法使用 defer ，直接放在最后
		// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
		// 可以把它单独放在一个函数中，这里是个实验，就这样了
		fw.Close()
		fr.Close()
	}
	return nil
}
