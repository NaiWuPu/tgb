package main

import (
	"fmt"
	"math/rand"
)

func main() {
	f1()
}

func f1() {
	rand.Seed(10)
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(20)
		fmt.Println(r1, r2)
	}
}
