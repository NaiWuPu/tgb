package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(strings.TrimRight("abba", "ba"))
	log.Println(strings.TrimSuffix("abcddcba", "dcba"))
	log.Println(strings.Split("qweqewqweqwewqqweqw1zklcnkzlnclcnl", "1"))
}
