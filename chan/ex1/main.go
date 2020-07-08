package main

import "fmt"

func main() {
	ch1 := make(chan bool, 2)
	ch1 <- true
	ch1 <- true

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

}
