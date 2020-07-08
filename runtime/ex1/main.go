package main

import (
	"fmt"
	"runtime"
)

// runtime.Caller
func f1()  {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}

	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(line)
}

func main() {
	f1()
}