package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime/pprof"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:9999", nil)
	}()
}

func CPU() {
	file, err := os.Create("./cpu.pprof")
	if err != nil {
		fmt.Printf("create cpu pprof failed, err :%v \n", err)
		return
	}
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	file.Close()
}

func MEM() {
	file, err := os.Create("./mem.pprof")
	if err != nil {
		fmt.Printf("create mem pprof failed, err: %v \n", err)
		return
	}
	pprof.WriteHeapProfile(file)
	file.Close()
}
