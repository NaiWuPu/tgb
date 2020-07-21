package main

import "fmt"

var target = []byte{0x50, 0x4b, 0x05, 0x06}

// 1010000 1001011 0000101 0000110
func main() {
	for _, v := range target {
		fmt.Println(string(v))
	}
}
