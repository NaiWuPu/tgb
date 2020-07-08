package main

import (
	"fmt"
	"time"
)

func f2() {
	now := time.Now()
	fmt.Println("now :",now)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-06-24 09:45:50", loc)
	fmt.Println(timeObj)
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	f2()
}
