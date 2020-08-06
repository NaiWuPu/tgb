package main

import (
	"bytes"
	"fmt"
)

func main() {
	fruit := []byte("apple")
	myFruit := fruit[2:3:4]
	fmt.Printf("myFruit %v \n", myFruit)
	for i, v := range myFruit {
		fmt.Printf("i:%v \n", i)
		fmt.Printf("v %v \n", v)
	}
	myFruit1 := fruit[2:3]
	fmt.Printf("myFruit1 %v \n", myFruit1)

	myFruit2 := fruit[3:4]
	fmt.Printf("myFruit2 %v \n", myFruit2)

	myFruit3 := fruit[2:4]
	fmt.Printf("myFruit3 %v \n", myFruit3)

	fmt.Println(fruit[:])
	fmt.Println(bytes.Equal(fruit, fruit[:]))
}
