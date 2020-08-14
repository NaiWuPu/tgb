package main

/*
#include "test.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(111)
	C.test()
}
