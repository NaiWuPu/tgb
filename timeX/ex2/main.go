package main

import (
	"fmt"
	"time"
)

const TIMEFORMAT = "2006-01-02 15:04:05"

func f2() {
	now := time.Now()

	fmt.Println("now :", now)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeObj, _ := time.ParseInLocation(TIMEFORMAT, "2020-08-14 12:00:00", loc)
	fmt.Println(timeObj)
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	f2()
}
