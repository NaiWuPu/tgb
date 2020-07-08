package main

import "fmt"

var a = []int64{1, 2, 3, 4, 5}
var b = []int64{1, 2}
func main() {
	for _, v := range b {
		bool := Find(a, v)
		fmt.Println(bool)
	}
}

// Find获取一个切片并在其中查找元素。如果找到它，它将返回它的密钥，否则它将返回-1和一个错误的bool。
func Find(slice []int64, val int64) ( bool) {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
