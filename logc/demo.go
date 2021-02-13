package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//filename := "./test.log"
	filename := "c:/logs/server.2021-01-04.log"
	start := time.Now()
	//Test_FileWrite()
	ReadString(filename) //17.57s 19.13s
	//ReadBytes(filename) //17.84s
	//ReadRune(filename) //20.88s
	//ReadLine(filename) //20.88s
	end := time.Now()
	fmt.Printf("readstring:%v\n", end.Sub(start))
}

func Test_FileWrite() {
	dstFile, err := os.OpenFile("c:/logs/demo.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	st := time.Now()
	defer func() {
		bufWriter.Flush()
		dstFile.Close()
		fmt.Println("文件写入耗时：", time.Now().Sub(st).Seconds(), "s")
	}()

	for i := 0; i < 100000; i++ {
		bufWriter.WriteString(strconv.Itoa(i) + "\n")
	}
}

func WriteString(str string) {
	dstFile, err := os.OpenFile("c:/logs/demo.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	bufWriter.WriteString(str + "\n")

	bufWriter.Flush()
	dstFile.Close()
}

func ReadString(filename string) {
	startTime := time.Now()
	defer func(startTime time.Time) {
		fmt.Printf("共耗时：%s\n", time.Now().Sub(startTime).String())
	}(startTime)
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	lineBegin := 0
	lineFirst := 0
	lineOver := 0
	//首次便利记录出现行号
	for {
		str, err := r.ReadString('\n')
		lineBegin = lineBegin + 1
		if err != nil {
			break
		}
		if strings.Contains(str, "1f83dcae3cfeebec") {
			fmt.Println(str)
			lineFirst = lineBegin
			lineBegin = lineBegin - 1000
			lineOver = lineFirst + 1000
			break
		}
	}
	fmt.Printf("查找耗时：%s\n", time.Now().Sub(startTime).String())
	/*fmt.Printf("lineBegin:%d\n", lineBegin)
	  fmt.Printf("lineOver:%d\n", lineOver)
	  fmt.Printf("lineFirst:%d\n", lineFirst)*/
	//二次遍历输出文件
	startTimeTwo := time.Now()
	start := 0
	dstFile, err := os.OpenFile("c:/logs/demo.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	for {
		str, err := r.ReadString('\n')
		start = start + 1
		if err != nil {
			break
		}
		if start >= lineBegin && start <= lineOver {
			bufWriter.WriteString(str)
		}
	}
	bufWriter.Flush()
	dstFile.Close()
	fmt.Printf("写出耗时：%s\n", time.Now().Sub(startTimeTwo).String())
}

func ReadBytes(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		_, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
	}
}

func ReadRune(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		c, size, err := r.ReadRune()
		if err != nil {
			break
		}
		fmt.Printf("%c %v\n", c, size)
	}
}

func ReadLine(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		_, err := readLine(r)
		if err != nil {
			break
		}
	}

}
func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
		if strings.Contains(string(line), "1f83dcae3cfeebec") {
			fmt.Println(string(line))
		}
	}
	return string(line), err
}
