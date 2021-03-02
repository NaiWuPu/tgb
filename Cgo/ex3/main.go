package main

/*
#include <string.h>
#include <stdlib.h>
#include <stdio.h>

void my_reverse(char* src, int len, char *dst){
  dst = malloc(sizeof(char) * (len + 1));
  printf("[c-part] src=%s", src);
  for (int i = 0; i < len; ++i)
  {
    dst[i] = src[len - 1 - i];
  }
  dst[len] = 0;
  printf("[c-part] dst=%s", dst);
}

void some_text(char* buffer, unsigned long long int *year){
  buffer = malloc(200 * sizeof(char));
  sscanf("year 2018d", "%s %16llu", buffer, year);
  printf("will return (%s, %16llu)", buffer, *year);

}
*/
import "C"

import "unsafe"
import "fmt"

func Reverse(a string) (dst string) {

	c_src := C.CString(a)
	defer C.free(unsafe.Pointer(c_src))
	c_len := C.int(len(a))
	c_dst := C.CString(dst)
	defer C.free(unsafe.Pointer(c_dst))

	C.my_reverse(c_src, c_len, c_dst)

	return string(*c_dst)

}

func Sometext() (dst string, year int64) {

	c_dst := C.CString("")
	c_year := C.ulonglong(0)
	defer C.free(unsafe.Pointer(c_dst))

	C.some_text(c_dst, &c_year)

	return string(*c_dst), int64(c_year)

}

func main() {
	fmt.Printf("[gopart] dst=%v", Reverse("Hello World"))

	buf, year := Sometext()
	fmt.Printf("received (%v, %v)", buf, year)
}
