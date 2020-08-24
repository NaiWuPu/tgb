package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var m2 = sync.Map{}
var wg = sync.WaitGroup{}
var x int64

func f1() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go f1()
	}
	wg.Wait()
	fmt.Println(x)
}
