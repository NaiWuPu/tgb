package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

var text = `60 104 116 109 108 62 10 60 109 101 116 97 32 104 116 116 112 45 101 113 117 105 118 61 34 114 101 102 114 101 115 104 34 32 99 111 110 116 101 110 116 61 34 48 59 117 114 108 61 104 116 116 112 58 47 47 119 119 119 46 98 97 105 100 117 46 99 111 109 47 34 62 10 60 47 104 116 109 108 62 10`

func main() {
	//buf := bytes.NewBufferString(text)
	//fmt.Println(BytesToString(buf))
	var t string
	for _, v := range strings.Split(text, " ") {
		z, _ := strconv.Atoi(v)
		t += string(rune(z))
	}
	fmt.Println(t)

}

func BytesToString(b io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(b)
	return buf.String()
}


func BytesToString2(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}