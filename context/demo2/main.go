package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)


func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
}

func f()  {
	defer wg.Done()
	for {
		fmt.Println("钱钱钱钱钱")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break
		default:

		}
	}
}