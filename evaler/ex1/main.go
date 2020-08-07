package main

import (
	"fmt"
	"github.com/soniah/evaler"
)

var In = make(chan int64, 50)

func main() {
	result, _ := evaler.Eval("(1+2.2+4.1*2)/3-0.31")
	res := evaler.BigratToFloat(result)
	res1, err := evaler.BigratToInt(result)
	if err != nil {
		fmt.Println(err)
	}
	res2 := evaler.BigratToBigint(result)
	fmt.Printf("运算结果1为:%v \n", result)
	fmt.Printf("运算结果2为:%v \n", res)
	fmt.Printf("运算结果3为:%v \n", res1)
	fmt.Printf("运算结果4为:%v \n", res2)

	go func() {
		for a := 0;; a++ {
			In <- 1
		}
	}()
	go func() {
		for {
			fmt.Println(<- In)
		}
	}()
	select {}
}
