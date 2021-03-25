package logs

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	models "logs/models"
	"os"

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

//Console 控制台
func (this *LogsController) Index() {
	//this.Ctx.WriteString("这是正则路由 user/test")
	this.TplName = "logs.html"
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
			req.ToFile(temppath + "/" + logs.Key + "/" + client.Ip + ".zip")
		} else {
			//获取所有客户端遍历
			clients, _ := models.QueryAllClient()
			for _, client := range clients {
				// 根据客户端ID查在线的客户端服务IP:Port
				url := "http://" + client.Ip + ":" + client.Port + "/file/query"
				// url := "http://localhost:2020/file/query"
				// 根据项目ID查日志路径、前缀、后缀，拼接日志全路径
				item := models.ReadItem(logs.Item)
				path := item.LogPath + item.LogPrefix + logs.Date + item.LogSuffix + ".log"
				// path := "C:\\logs\\server.2021-01-04.log"
				//通过Http调用客户端
				// go getLog(url, path, logs.Key, temppath, client.Ip, logs.Line)
				req := httplib.Post(url).Debug(true)
				req.JSONBody(map[string]interface{}{"path": path, "key": logs.Key, "line": logs.Line})
				req.ToFile(temppath + "/" + logs.Key + "/" + client.Ip + ".zip")
			}
		}
		Zip(temppath+"/"+logs.Key+".zip", temppath+"/"+logs.Key+"/")
		// this.Ctx.WriteString("好了")
		this.Ctx.Output.Download(temppath + "/" + logs.Key + ".zip")
	}
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
