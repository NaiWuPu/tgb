package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)
var i int

func main() {
	go f1()
	go f2()
	select {}
}

func f1() {
	for {
		fmt.Println("外层for循环")
		select {
		case chLog := <-ch:
			if i == 0 {
				i = 1
				fmt.Println("内层循环 跳了出去")
				continue
			}
			if i == 1 {
				i = 0
			}
			fmt.Println("time", chLog)
		default:
			fmt.Println("我已沉睡")
			time.Sleep(2 * time.Second)
		}
	}
}

func f2() {
	for {
		ch <- time.Now().Second()
		time.Sleep(500 * time.Microsecond)
	}
}
