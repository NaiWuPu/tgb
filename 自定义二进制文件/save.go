package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website struct {
	Url int32
}

func main() {
	file, err := os.Create("output.bin")
	for i := 1; i <= 10; i++ {
		info := Website{
			int32(i),
		}
		if err != nil {
			fmt.Println("文件创建失败 ", err.Error())
			return
		}
		defer file.Close()
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}
