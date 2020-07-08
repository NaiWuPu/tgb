package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	LOOP:
		for {
			fmt.Printf("worker, trace code:%s\n", traceCode)
			time.Sleep(time.Millisecond * 10) // 数据库链接时间
			select {
			case <-ctx.Done():
				break LOOP
			default:
			}
		}
	fmt.Println("worder done!")
	wg.Done()
}

func main() {
	// 设置一个50 毫秒超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)
	// 在系统的入口中设置trace code 传递给后续启动的goroutine 实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "1251231234")
	wg.Add(1)

	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}