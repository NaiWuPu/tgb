package main

import "fmt"

func main() {

	f1()
}

func f1() (r []byte) {
	r = make([]byte, 1024)
	for i := 0; i < 1000; i++ {
		r = append(r, byte(i))
	}
	fmt.Println(string(r))
	return r
}
