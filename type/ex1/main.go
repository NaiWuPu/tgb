package main

import (
	"fmt"
)

type Hello struct {
	A int `json:"a"`
}

func main() {
	B := new(Hello)
	D := Hello{A: 1}
	fmt.Println()
	fmt.Printf("%v \n", &B)
	fmt.Printf("%v \n", &D)

}
