package main

import (
	"fmt"
	"os"
)

func main() {
	f1()
	fmt.Println("--")
	f2()
}

func f1() {
	f1, _ := os.Stat("demo")
	f2, _ := os.Stat("demo")

	fmt.Println(f1.Name())
	fmt.Println(f1.IsDir())
	fmt.Println(f1.ModTime())
	fmt.Println(f1.Mode())
	fmt.Println(f1.Size())
	fmt.Println(f1.Sys())
	if os.SameFile(f1, f2) {
		fmt.Println("same!")
		return
	}
	fmt.Println("not same!")
}

func f2() {
	f1, _ := os.Lstat("demo")
	fmt.Println(f1.Name())
	fmt.Println(f1.IsDir())
	fmt.Println(f1.ModTime())
	fmt.Println(f1.Mode())
	fmt.Println(f1.Size())
	fmt.Println(f1.Sys())
}
