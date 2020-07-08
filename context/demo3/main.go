package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

func f(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("钱钱钱钱钱")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}
