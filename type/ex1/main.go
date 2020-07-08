package main

import "fmt"

func main() {
	var x = [3]int{1,2,3}
	fmt.Println(x)
	a(&x)
	fmt.Println(x)
}

func a(a *[3]int)  {
	a[1] = 100
	fmt.Println(a)
}