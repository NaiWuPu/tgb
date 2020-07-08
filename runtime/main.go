package main

import (
	"fmt"
	"runtime"
)
func init() {
	runtime.GOMAXPROCS(1)
}
func main() {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("archive:", runtime.GOOS)
}