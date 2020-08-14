package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	v := 42
	fmt.Println(unsafe.Pointer(&v))
	C.printint(C.int(unsafe.Pointer(&v)))
	return
}
