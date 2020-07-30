package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(time.Second * 2)

	ch := make(chan bool)
	go func(t *time.Timer) {
		defer t.Stop()
		for {
			select {
			case <-t.C:
				fmt.Println("timer running....")
				// 需要重置Reset 使 t 重新开始计时
				t.Reset(time.Second * 2)
			case stop := <-ch:
				if stop {
					fmt.Println("timer Stop")
					return
				}
			}
		}
	}(t)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}
