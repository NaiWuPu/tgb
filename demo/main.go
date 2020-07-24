package main

import (
	"fmt"
	"sync"
)

var sy1 sync.Mutex
var sy2 sync.Mutex

var wg sync.WaitGroup
var a int
var b int

func main() {

	go f1()
	go f1()
	select {}
}

func f1() {
	for {
		if a != 0 || b != 0 {
			fmt.Println(a, b)
			fmt.Println("没锁住")
			continue
		}
		sy1.Lock()
		a++
		sy2.Lock()
		b++
		b--
		sy2.Unlock()
		a--
		sy1.Unlock()
	}
}
