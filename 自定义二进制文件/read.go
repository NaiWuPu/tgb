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
	file, err := os.Open("output.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	m := Website{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}
func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}
