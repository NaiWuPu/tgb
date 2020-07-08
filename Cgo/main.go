package main


import "C"


import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	//使用C.CString创建的字符串需要手动释放。
	cs := C.CString("Hello World\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
	fmt.Println("call C.sleep for 3s")
	C.sleep(3)

	runtime.Caller()
	return
}