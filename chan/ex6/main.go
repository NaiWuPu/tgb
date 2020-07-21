package main

import (
	"fmt"
	"time"
)

var ch = make(chan int64, 1)

func main() {
	go f2()
	go f1()
	select {}
}

func f1() {
	for {
		select {
		case t := <-ch:
			fmt.Printf("输入了%d\n", t)
		default:
			fmt.Println("sleep")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func f2() {
	for {
		ch <- time.Now().Unix()
		time.Sleep(1 * time.Second)
	}
}
