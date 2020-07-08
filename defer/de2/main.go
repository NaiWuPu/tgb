package main

import "fmt"

func f4() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}


func main()  {

	fmt.Println(f4())
}