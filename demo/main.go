package main

import (
	"fmt"
)

func main() {
	err := e()
	cher(err)

	if nil == nil {
		fmt.Println(3333)
	}
}

func e() (err error) {
	return nil
}

func cher(err error) {
	if err != nil {
		fmt.Println("111")
		return
	}
	fmt.Println("2222")
}
