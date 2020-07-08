package main

import "fmt"

func f(arg int64) {
	fmt.Println( arg)
}

func main()  {
	f(1|2|3|4|5|6|6|7)
}
