package main

import (
	"fmt"
	"time"
)

func f1(ch chan int)  {
	for {
		ch <- 1
	}
}

func f2(ch chan int)  {
	for {
		ch <- 2
	}
}

func main() {
	ch := make(chan int, 20)
	fmt.Println(<-ch)

	go f1(ch)

	go f2(ch)

	for {
		fmt.Println(<-ch)

		time.Sleep(1 * time.Second)
	}

}