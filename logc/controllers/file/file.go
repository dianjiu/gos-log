package file

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
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
		defer func() {
			//返回后清理压缩文件
			os.Remove(file)
		}()
		this.Ctx.Output.Download(file)
	}
}

//ReadString
func ReadString(filename string, key string, line int64) (file string) {
	startTime := time.Now()
	defer func(startTime time.Time) {
		log.Println("共耗时：%s\n", time.Now().Sub(startTime).String())
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
		log.Println("open file failed, err:%v", err)
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
			// log.Println(str)
			if lineFirst == 0 && lineOver == 0 {
				lineFirst = lineBegin
				lineOver = lineFirst + line
				log.Printf("lineFirst:%d\n", lineFirst)
				log.Printf("lineBegin:%d\n", lineBegin)
				log.Printf("lineOver:%d\n", lineOver)
			}
		}
		// 暂不考虑lineOver 小于 over 的时候
		if lineBegin < lineOver {
			bufWriter.WriteString(str)
		}
		if lineBegin == lineOver {
			bufWriter.WriteString(str)
			//输出换行
			bufWriter.WriteString("\n")
		}
		if lineOver > 0 && lineBegin > lineOver {
			//break
			// 六一儿童节预览版
			if strings.Contains(str, key) {
				// log.Println(str)
				if lineFirst != 0 && lineOver != 0 {
					lineFirst = lineBegin
					lineOver = lineFirst + line
					log.Printf("lineFirst:%d\n", lineFirst)
					log.Printf("lineBegin:%d\n", lineBegin)
					log.Printf("lineOver:%d\n", lineOver)
				}
			}
		}
	}
	bufWriter.Flush()
	dstFile.Close()
	log.Println("查找耗时：%s\n", time.Now().Sub(startTime).String())
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
	log.Println("写出耗时：%s\n", time.Now().Sub(startTimeTwo).String()) */
	// 压缩
	startTimeThree := time.Now()
	ip := GetLocalIPv4()
	var src string = temppath + key + ".log"
	var dst string = temppath + ip + ".zip"
	Zip(dst, src)
	log.Println("压缩耗时：%s\n", time.Now().Sub(startTimeThree).String())
	return dst
}

//GetLocalIPv4 获取本机的IPv4地址
func GetLocalIPv4() (ip string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Println("net.Interfaces failed, err:", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						log.Println("本机IP：" + ipnet.IP.String())
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
		log.Println("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)
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
		log.Println("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

		// 因为是在循环中，无法使用 defer ，直接放在最后
		// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
		// 可以把它单独放在一个函数中，这里是个实验，就这样了
		fw.Close()
		fr.Close()
	}
	return nil
}
