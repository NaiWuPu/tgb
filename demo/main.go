package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	fmt.Printf("%v\n", ptr)
}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := range x {
		x[i] = i
	}
	return x
}
