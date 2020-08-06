package main

import (
	"compress/gzip"
	"log"
	"os"
	"strings"
)

func main() {

	// 获取要打包的文件信息
	fr, err := os.Open("demo.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 获取文件头信息
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(fi.Name())
	//fmt.Println(fi.Mode())
	//fmt.Println(fi.IsDir())
	//fmt.Println(fi.Sys())

	fileName := strings.Split(fi.Name(), ".")[0]

	fw, err := os.Create(fileName + ".gzip") // 创建gzip包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	// 实例化心得gzip.Writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 创建gzip.Header
	gw.Header.Name = fi.Name()

	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// 写入数据到zip包
	_, err = gw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
}
